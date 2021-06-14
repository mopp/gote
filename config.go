package main

import "os"

type Config struct {
	note_dir string
}

func newConfig() Config {
	return Config{
		note_dir: os.Getenv("HOME") + "/notes/",
	}
}
