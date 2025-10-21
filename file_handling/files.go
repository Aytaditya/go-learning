package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, error := os.Open("example.txt")
	if error != nil {
		panic(error) // panic is used here for simplicity; in production code, handle errors appropriately
	}
	fileInfo, fileError := f.Stat()
	if fileError != nil {
		panic(fileError)
	}
	defer f.Close()

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())    // prints the size of the file in bytes
	fmt.Println(fileInfo.IsDir())   // prints false as example.txt is a file not a folder
	fmt.Println(fileInfo.ModTime()) // prints the last modified time of the file

	// reading content of the file
	buffer := make([]byte, 20) // buffer is just an array of bytes
	_, err := f.Read(buffer)   // Read function reads up to len(buffer) bytes from the file
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buffer); i++ {
		fmt.Print(string(buffer[i])) // converting byte to string
	}
	fmt.Println()

	// second way of reading file (Easiest way)
	filetext, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(filetext))

	// READ FOLDERSS
	data, err := os.ReadDir("../")
	if err != nil {
		panic(err)
	}
	fmt.Print(data)

	// CREATING A FILE
	newFile, err := os.Create("newfile.txt")
	if err != nil {
		panic(err)
	}

	defer newFile.Close()

	newFile.WriteString("Hello World")
	newFile.WriteString(" This is a new file. ") // this will append to the file not replace

	// Transferring data from one file to another
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	destFile, err := os.Create("destfile.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte() // return one byte at a time and error
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		er := writer.WriteByte(b)
		if err != nil {
			panic(er)
		}
	}

	writer.Flush()
	fmt.Println()
	fmt.Println("File copied successfully")

}
