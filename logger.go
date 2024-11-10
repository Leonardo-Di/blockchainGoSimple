package main

import (
	"log"
	"os"
)

var (
	logger *log.Logger
)

func init() {
	file, err := os.OpenFile("blockchain.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(message string) {
	logger.Println(message)
}

func LogError(err error) {
	logger.Println("ERROR:", err)
}
