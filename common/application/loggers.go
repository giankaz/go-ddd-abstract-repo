package common

import (
	"log"
	"os"
)

var ErrorLog *log.Logger

func NewErrorLog() {

	newLog := log.New(os.Stdout, "[MONGODB - ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = newLog
}

var RoutineLog *log.Logger

func NewRoutineLog() {

	newLog := log.New(os.Stdout, "[MONGODB - ROUTINE]: ", log.Ldate|log.Ltime|log.Lshortfile)
	RoutineLog = newLog
}

var InfoLog *log.Logger

func NewInfoLog() {

	newLog := log.New(os.Stdout, "[MONGODB - INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLog = newLog
}