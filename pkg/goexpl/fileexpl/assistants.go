package fileexpl

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	filename1 = "./files/sample.txt"
	filew1    = "./files/dat1"
	filew2    = "./files/dat2"
	subdir    = "./files/subdir"
)

var pl = fmt.Println

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

func printFilesStructure() {
	printTree("./files", "")
}

func printTree(root string, prefix string) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for i, entry := range entries {
		// 确定是否是最后一个条目
		isLast := i == len(entries)-1

		// 打印当前条目
		fmt.Print(prefix)
		if isLast {
			fmt.Print("└── ")
		} else {
			fmt.Print("├── ")
		}
		fmt.Println(entry.Name())

		// 如果是目录，递归打印其内容
		if entry.IsDir() {
			newPrefix := prefix
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}

			err := printTree(filepath.Join(root, entry.Name()), newPrefix)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	pl(" ", path, d.IsDir())
	return nil
}
