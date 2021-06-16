package app

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Note struct {
	dir  string
	name string
	text string
}

func newNote(dir string, name string) *Note {
	return &Note{
		dir:  dir,
		name: name,
	}
}

func (n *Note) String() string {
	return n.name
}

func (n *Note) SetText(text string) {
	n.text = text
}

func (n *Note) Read() (string, error) {
	text, err := ioutil.ReadFile(n.path())

	if err != nil {
		return "", err
	}

	return string(text), nil
}

func (n *Note) save() error {
	f, err := os.OpenFile(n.path(), os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	_, err = io.Copy(f, strings.NewReader(n.text))

	return err
}

func (n *Note) path() string {
	return n.dir + n.name
}
