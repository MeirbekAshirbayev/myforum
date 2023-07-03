package main

import (
	"forum/server"
	"log"
	"os"
)

type Logger struct {
	log *log.Logger
}

func main() {
	// Create the log file
	file, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a custom logger
	logger := Logger{log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)}

	// Start your server and pass the logger
	server.NewServer(logger.log)
}
