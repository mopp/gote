package gote

import "os"

type Config struct {
	noteDir              string
	logFilepath          string
	dailyNoteTitleFormat string
}

func NewConfig() *Config {
	return &Config{
		noteDir:              os.Getenv("HOME") + "/notes/",
		logFilepath:          "gote.log",
		dailyNoteTitleFormat: "2006-01-02",
	}
}
