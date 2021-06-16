package app

import "os"

type Config struct {
	noteDir string
	logFilepath string
}

func NewConfig() *Config {
	return &Config{
		noteDir: os.Getenv("HOME") + "/notes/",
		logFilepath: "gote.log",
	}
}
