package main

import (
	"fmt"
	"os"
)

func main() {
	i := 1
	for {
		generateFile(i)
		i++
	}
}

func generateFile(i int) {
	fileName := fmt.Sprintf("./tmp/file%d.txt", i)
	fileContent := fmt.Sprintf("This is file %d", i)

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(fileContent)
}
