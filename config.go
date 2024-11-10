package main

import (
	"os"
	"strconv"
)

type Config struct {
	Port       string
	Difficulty int
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	difficulty := 4 // Default difficulty
	if d := os.Getenv("DIFFICULTY"); d != "" {
		if parsed, err := strconv.Atoi(d); err == nil {
			difficulty = parsed
		}
	}

	return &Config{
		Port:       port,
		Difficulty: difficulty,
	}
}
