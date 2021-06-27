package gote

import (
	"fmt"
	"io/ioutil"
	"time"
)

type Service struct {
	config *Config
}

func NewService(c *Config) *Service {
	return &Service{
		config: c,
	}
}

func (s *Service) CreateNote(name string) (*Note, error) {
	n := newNote(s.config.noteDir, name)
	err := n.Write("")
	if err != nil {
		return nil, fmt.Errorf("could not create note: %w", err)
	}

	return n, nil
}

func (s *Service) FindDailyNoteToday() (*Note, error) {
	today := s.generateDailyNoteBasenameToday()
	title := today + ".md"

	n, err := s.findBy(title)
	if err != nil {
		return nil, err
	}

	return n, nil
}

func (s *Service) CreateDailyNoteTody() (*Note, error) {
	today := s.generateDailyNoteBasenameToday()
	title := today + ".md"

	n, err := s.findBy(title)
	if err != nil {
		// TODO: Define error struct.
		return nil, fmt.Errorf("already exist: %s", title)
	}

	if n == nil {
		n = newNote(s.config.noteDir, title)
		t :=
			`## %s

### 今日やったこと

### 明日やること

### 雑記

`

		err := n.Write(fmt.Sprintf(t, today))
		if err != nil {
			return nil, err
		}
	}

	return n, nil
}

func (s *Service) FetchAllNotes() ([]*Note, error) {
	files, err := ioutil.ReadDir(s.config.noteDir)
	if err != nil {
		return nil, fmt.Errorf("could not fetch notes at %s: %w", s.config.noteDir, err)
	}

	notes := make([]*Note, len(files))
	for i, f := range files {
		notes[i] = newNote(s.config.noteDir, f.Name())
	}

	return notes, nil
}

func (s *Service) findBy(title string) (*Note, error) {
	notes, err := s.FetchAllNotes()
	if err != nil {
		return nil, err
	}

	for _, n := range notes {
		if n.Title() == title {
			return n, nil
		}
	}

	return nil, nil
}

func (s *Service) generateDailyNoteBasenameToday() string {
	return time.Now().Format(s.config.dailyNoteTitleFormat)
}
