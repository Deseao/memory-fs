package main

import (
	"fmt"

	"github.com/Deseao/memory-fs/internal/fs"
)

func main() {
	root := fs.New("root")
	root.CreateDirectory([]string{"poems"})
	root.CreateDirectory([]string{"poems", "love"})
	shortLovePoem := fs.NewFile("roses.txt", []byte("I miss you ma."))
	otherFile := fs.NewFile("surpriseBirthdayParty.png", []byte("Image data gibberish"))
	root.CreateDirectory([]string{"novellas"})
	fmt.Println(root.Print(0))
	fmt.Println(shortLovePoem.Name)
	fmt.Println(string(shortLovePoem.Data))
	root.Save([]*fs.File{shortLovePoem, otherFile}, []string{"poems", "love", "letters"})
	poems := root.Subdirectories[0]
	love := poems.Subdirectories[0]
	letters := love.Subdirectories[0]
	fmt.Println("Letters name: ", letters.Name)
	fmt.Println("Files length: ", len(letters.Files))
	for _, f := range letters.Files {
		fmt.Println(f.Name, string(f.Data))
	}
}
