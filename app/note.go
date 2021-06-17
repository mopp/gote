package app

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Note struct {
	dir  string
	name string
}

func newNote(dir string, name string) *Note {
	return &Note{
		dir:  dir,
		name: name,
	}
}

func (n *Note) Title() string {
	return n.name
}

func (n *Note) Read() (string, error) {
	text, err := ioutil.ReadFile(n.path())

	if err != nil {
		return "", err
	}

	return string(text), nil
}

func (n *Note) Write(t string) error {
	return n.WriteFrom(strings.NewReader(t))
}

func (n *Note) WriteFrom(w io.Reader) error {
	f, err := os.OpenFile(n.path(), os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return fmt.Errorf("could not open note at %s: %w", n.path(), err)
	}

	_, err = io.Copy(f, w)

	return err
}

func (n *Note) String() string {
	return n.Title()
}

func (n *Note) path() string {
	return n.dir + n.name
}
