package service

import (
	"errors"
	"forum/entity"
	"net/mail"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) Signup(user *entity.User) error {
	b, err := s.repo.CheckUsername(user.Username)
	if err != nil || b {
		return err
	}
	b, err = s.repo.CheckEmail(user.Email)
	if err != nil || b {
		return err
	}
	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		return err
	}
	if !isPasswordValid(user.Password) {
		return errors.New("Invalid password. Password must be at least 8 characters long, contain capitalized alphabets, and special symbols.")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.repo.CreateUser(user)
}

func isPasswordValid(password string) bool {
	// Check minimum length requirement
	if len(password) < 8 {
		return false
	}

	// Check for at least one capitalized alphabet
	hasCapital := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasCapital {
		return false
	}

	// Check for at least one special symbol
	hasSpecialSymbol := regexp.MustCompile(`[!@#$%^&*()]`).MatchString(password)
	if !hasSpecialSymbol {
		return false
	}

	return true
}

func (s *service) Login(email string, password string) (*entity.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
