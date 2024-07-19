package main

import (
	"fmt"
	"os"
)

type Writer interface {
	Write([]byte) (int, error)
}

type File struct {
	file *os.File
}

func NewFile(filename string) (*File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return &File{file: file}, nil
}

func (f *File) Write(text []byte) (int, error) {
	return f.file.Write(text)
}

func (f *File) Close() error {
	return f.file.Close()
}

type Console struct {}

func (c *Console) Write(text []byte) (int, error) {
	return fmt.Printf(string(text))
}

func main() {
	file, err := NewFile("output.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла: ", err)
		return
	}
	defer file.Close()


	console := Console{}

	text := "Hello, world!"

	_, err = file.Write([]byte(text))
	if err != nil {
		fmt.Println("Ошибка записи в файл: ", err)
		return
	}


	_, err = console.Write([]byte(text))
	if err != nil {
		fmt.Println("Ошибка записи в консоль: ", err)
		return
	}
}