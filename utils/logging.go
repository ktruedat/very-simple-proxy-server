package utils

import (
	"log"
	"os"
)

var (
	ErrLogger  *log.Logger
	InfoLogger *log.Logger
)

func NewInfoLogger() *log.Logger {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return InfoLogger
}

func NewErrLogger() *log.Logger {
	ErrLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return ErrLogger
}
