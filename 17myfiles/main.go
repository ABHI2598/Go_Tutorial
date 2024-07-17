package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "Hello my name is Abhinav Pandey.This is sample txt file for files"

	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Total length written in file: %d\n", length)
	defer file.Close()
	ReadFiles("./myfile.txt")
	ReadFilesUsinfBuffer()

}

func ReadFiles(file string) {
	databytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))
}

func ReadFilesUsinfBuffer() {

	file, err := os.Open("./myfile.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	buffer := make([]byte, 1024)

	for {

		content, err := file.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(string(buffer[:content]))
	}

}

func readFile() {

	file, err := os.Open("./myfile.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	buf := make([]byte, 1024)

	for {

		content, err := file.Read(buf)

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Printf(string(buf[:content]))

	}

}
