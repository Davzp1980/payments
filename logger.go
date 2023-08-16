package main

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	fileInfo, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fileWarning, err := os.OpenFile("logsWarning.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fileError, err := os.OpenFile("logsError.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(fileInfo, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(fileWarning, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(fileError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
