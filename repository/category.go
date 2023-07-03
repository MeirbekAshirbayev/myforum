package repository

func (r repo) CreateCategory(names []string, postID int) error {
	for _, name := range names {
		_, err := r.db.Exec("INSERT INTO Category (PostID, Name) VALUES (?, ?)", postID, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r repo) DeleteCategoryByPostID(postID int) error {
	_, err := r.db.Exec("DELETE FROM Category WHERE PostID = ?", postID)
	if err != nil {
		return err
	}
	return nil
}
