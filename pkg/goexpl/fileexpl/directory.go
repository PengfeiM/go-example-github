// Package fileexpl file system example
package fileexpl

import (
	"fmt"
	"os"
	"path/filepath"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// DirExpl Go has several userful function for working with directories in the FS
type DirExpl struct{}

func (de *DirExpl) RunExample(inputParams *goexpl.InputParams) error {
	err := os.Mkdir(subdir, 0o755)
	check(err)
	defer os.RemoveAll(subdir)
	pl("create a subdir in files")
	printFilesStructure()

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0o644))
	}

	createEmptyFile(fmt.Sprintf("%s/file1", subdir))
	pl("create a file in subdir")
	printFilesStructure()

	// may be we can create dirs with many layers
	err = os.MkdirAll(fmt.Sprintf("%s/sub1/sub2/sub3", subdir), 0o755)
	check(err)
	pl("create hierarchy directories")
	printFilesStructure()

	createEmptyFile(fmt.Sprintf("%s/sub1/file2", subdir))
	createEmptyFile(fmt.Sprintf("%s/sub1/file3", subdir))
	createEmptyFile(fmt.Sprintf("%s/sub1/sub2/file4", subdir))
	pl("prepare a file tree")
	printFilesStructure()
	c, err := os.ReadDir(fmt.Sprintf("%s/sub1", subdir))
	check(err)
	pl("Listing subdir, you may compare the result with former output")
	for _, entry := range c {
		pl(" ", entry.Name(), entry.IsDir())
	}

	pl("we can do cd with Chdir")
	err = os.Chdir(fmt.Sprintf("%s/sub1/sub2/", subdir))
	check(err)
	c, err = os.ReadDir(".")
	check(err)
	pl("go into ./subdir/sub1/sub2/\n and let's see how it looks like")
	for _, entry := range c {
		pl(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../../..")
	check(err)

	pl("Visiting subdir")
	err = filepath.WalkDir(subdir, visit)
	check(err)

	return nil
}

func (de *DirExpl) Init() {
	if err := goexpl.RegisterGoExample(de.GetGoExampleName(), &DirExpl{}); err != nil {
		panic(err)
	}
}

func (de *DirExpl) GetGoExampleName() string {
	return common.DirExpl
}
