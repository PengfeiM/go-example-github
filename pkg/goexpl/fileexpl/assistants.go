package fileexpl

import (
	"fmt"
	"io"
	"os"
)

var (
	filename1 = "./files/sample.txt"
	filew1    = "./files/dat1"
	filew2    = "./files/dat2"
)

// check
// panic err if err isn't null
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func printFilePointer(f *os.File) {
	pos, err := f.Seek(0, io.SeekCurrent)
	check(err)
	buf := make([]byte, 1)
	_, err = f.Read(buf)
	check(err)
	_, err = f.Seek(-1, io.SeekCurrent)
	check(err)
	fmt.Printf("file pointer @ %d, char: %s\n", pos, string(buf))
}
