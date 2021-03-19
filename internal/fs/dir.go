package fs

import (
	"strings"
)

type Directory struct {
	Name           string
	Files          []*File
	Subdirectories []*Directory
}

func New(name string) *Directory {
	return &Directory{
		Name: name,
	}
}

func (d *Directory) CreateDirectory(path []string) {
	if len(path) == 0 {
		return
	}

	head := path[0]
	for _, subDir := range d.Subdirectories {
		if subDir.Name == head {
			subDir.CreateDirectory(path[1:])
			return
		}
	}
	newDir := New(head)
	d.Subdirectories = append(d.Subdirectories, newDir)
	newDir.CreateDirectory(path[1:])
}

func (d *Directory) Print(depth int) string {
	result := strings.Repeat("-", depth) + d.Name + "\n"
	for _, v := range d.Subdirectories {
		result += v.Print(depth + 1)
	}
	return result
}

func (d *Directory) Save(files []*File, path []string) {
	if len(path) == 0 {
		d.Files = append(d.Files, files...)
		return
	}
	head := path[0]
	for _, subDir := range d.Subdirectories {
		if subDir.Name == head {
			subDir.Save(files, path[1:])
			return
		}
	}
	newDir := New(head)
	d.Subdirectories = append(d.Subdirectories, newDir)
	newDir.Save(files, path[1:])
}
