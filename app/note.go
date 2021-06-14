package app

import (
	"io/fs"
	"io/ioutil"
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
	text, err := ioutil.ReadFile(n.dir + "/" + n.fi.Name())

	if err != nil {
		return "", err
	}

	return string(text), nil
}
