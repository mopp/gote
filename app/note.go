package app

import (
	"io"
	"io/ioutil"
	"os"
)

type Note struct {
	dir  string
	name string
}

func NewNote(dir string, name string) Note {
	return Note{
		dir:  dir,
		name: name,
	}
}

func (n *Note) String() string {
	return n.name
}

func (n *Note) Read() (string, error) {
	text, err := ioutil.ReadFile(n.path())

	if err != nil {
		return "", err
	}

	return string(text), nil
}

func (n *Note) Save(text io.Reader) error {
	f, err := os.OpenFile(n.path(), os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	_, err = io.Copy(f, text)

	return err
}

func (n *Note) path() string {
	return n.dir + n.name
}
