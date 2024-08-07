package main

import (
	"fmt"
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

var (
	fileInfo *os.FileInfo
	fileErr  error
)

func main() {

	// newFile, err = os.Create("demo.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fileInfo, fileErr := os.Stat("demo.txt")

	if fileErr != nil {
		log.Fatal(fileErr)
	}

	fmt.Println("File Name: ", fileInfo.Name())

	fmt.Println("Size in bytes: ", fileInfo.Size())

	fmt.Println("Permissions: ", fileInfo.Mode())

	fmt.Println("Last Modified:", fileInfo.ModTime())

	fmt.Println("Is Dictionary: ", fileInfo.IsDir())
}
