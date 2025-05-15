package fileexpl

import (
	"bufio"
	"fmt"
	"os"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// FileWriteExpl
// besides file reading, we can also write files
type FileWriteExpl struct{}

func (fwe *FileWriteExpl) RunExample(inputParams *goexpl.InputParams) error {
	// let's try to dump string([]byte) to file
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile(filew1, d1, 0o644)
	check(err)

	// or you can create a file and get the file pointer
	f, err := os.Create(filew2)
	check(err)
	defer f.Close() // never forget to close your file

	// write slices to file
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// write a string directly
	n3, err := f.WriteString("I write a string\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// the writings above can't affect the file on disk, you need f.Sync() to make it saved on disk
	f.Sync()

	// use bufio, we can not only read file, but also write files
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("I write file with bufio")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush() // flush all buffer to files

	return nil
}

func (fwe *FileWriteExpl) Init() {
	if err := goexpl.RegisterGoExample(fwe.GetGoExampleName(), &FileWriteExpl{}); err != nil {
		panic(err)
	}
}

func (fwe *FileWriteExpl) GetGoExampleName() string {
	return common.FileWriteExpl
}
