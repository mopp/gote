package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Service struct {
	config *Config
}

func NewService(c *Config) *Service {
	return &Service{
		config: c,
	}
}

func (s *Service) CreateNote(name string) *Note {
	n := newNote(s.config.noteDir, name)
	n.save()

	return n
}

func (s *Service) UpdateNote(n *Note) {
	n.save()
}

func (s *Service) FetchAllNotes() []*Note {
	files, err := ioutil.ReadDir(s.config.noteDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		log.Fatal(err)
	}

	notes := make([]*Note, len(files))
	for i, file := range files {
		notes[i] = newNote(s.config.noteDir, file.Name())
	}

	return notes
}
