package file

import (
	"log"
	"os"
)

func Create(filename string) {
	// Creates a new file in your directory.
	// Supported filetype: ???
	newFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err) // print + os.Exit()
	}
	log.Println(newFile) // TODO: Change this later
	newFile.Close()
}

func Delete(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}
}
