package main

import (
	"github.com/3nt3/homework/color"
	"io"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	mw := io.MultiWriter(file, os.Stdout)

	WarningLogger = log.New(mw, color.Yellow+"[WARNING] "+color.Reset, log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(mw, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(mw, color.Red+"[ERROR] "+color.Reset, log.Ldate|log.Ltime|log.Lshortfile)


	WarningLogger.Println("test")
	ErrorLogger.Println("test")
	InfoLogger.Println("test")
}
