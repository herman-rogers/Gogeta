package main

import (
	"log"
	"os"
	"time"
)

func GetLogFile() *os.File {
	var time = time.Now().Local()
	var filename string = "gogeta_" + time.Format("2006-01-02") + ".log"
	var directory string = "logs/" + filename
	logfile, err := os.OpenFile(directory, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Print("Open File Error: " + filename + " not found")
	}
	return logfile
}

func LoggerInfo(data string) {
	logfile := GetLogFile()
	defer logfile.Close()
	log.SetOutput(logfile)
	log.Print(data)
}

func LoggerError(data string) {
	logfile := GetLogFile()
	defer logfile.Close()
	log.SetOutput(logfile)
	log.Fatal(data)
}
