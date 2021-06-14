package app

import (
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

type Note struct {
	dir string
	fi  fs.FileInfo
}

func NewNote(dir string, fi fs.FileInfo) Note {
	return Note{
		dir: dir,
		fi:  fi,
	}
}

func (n *Note) String() string {
	return n.fi.Name()
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
	return n.dir + "/" + n.fi.Name()
}
