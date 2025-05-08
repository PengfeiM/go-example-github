package fileexpl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// FileReadExpl
// first, we read file
type FileReadExpl struct{}

func (fre *FileReadExpl) RunExample(inputParams *goexpl.InputParams) error {
	// just reading file
	dat, err := os.ReadFile(filename1)
	check(err)
	fmt.Println(string(dat))
	// by reading file, you can just operate file once.
	// if you wish to do it multi times, open it and you will get a file pointer
	f, err := os.Open(filename1)
	check(err)
	defer f.Close()
	printFilePointer(f)

	// use a []byte to read it
	fmt.Println("read 5 bytes from beginning")
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	printFilePointer(f)

	// jump to a known location in the file relative to the beginning
	fmt.Println("seek relative to beginning and read 2 bytes")
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))
	printFilePointer(f)

	// there are other ways to move file pointer in file.
	// relative to current position
	fmt.Println("move file pointer by -6 relative to current position")
	_, err = f.Seek(-6, io.SeekCurrent)
	check(err)
	printFilePointer(f)
	// relative to ending
	fmt.Println("move file pointer by -3 relative to ending")
	_, err = f.Seek(-3, io.SeekEnd)
	check(err)
	printFilePointer(f)

	// request for certain number of char
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	printFilePointer(f)

	// move f to beginning
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// bufio is a more powerful file operator
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	return nil
}

func (fre *FileReadExpl) Init() {
	if err := goexpl.RegisterGoExample(fre.GetGoExampleName(), &FileReadExpl{}); err != nil {
		panic(err)
	}
}

func (fre *FileReadExpl) GetGoExampleName() string {
	return common.FileReadExpl
}
