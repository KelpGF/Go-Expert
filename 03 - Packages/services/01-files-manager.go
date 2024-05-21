package services

import (
	"bufio"
	"fmt"
	"os"
)

const fileName = "test.txt"

func Run01() {
	createFile()
	readFileSync()
	readFileChunks()
	deleteFile()
}

func createFile() {
	// create a file
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	// size, err := f.WriteString("Hello, World!")
	size, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", size)

	f.Close()
}

// load all file content into memory
func readFileSync() {
	// open a file
	f, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(f))
}

// load file content in chunks
func readFileChunks() {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}

func deleteFile() {
	err := os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
}
