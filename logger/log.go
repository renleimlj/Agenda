package logger

import (
	"log"
	"os"
)

func Log(message string) {
	logFile, _ := os.OpenFile("Log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	defer logFile.Close()
	infoLog := log.New(logFile, "[Info]", log.LstdFlags)
	infoLog.Println(message)
}