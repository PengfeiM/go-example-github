package fileexpl

import (
	"path/filepath"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// FilePathExpl
// here we use path/filepath package to operate filepaths
type FilePathExpl struct{}

func (fp *FilePathExpl) RunExample(inputParams *goexpl.InputParams) error {
	pl("# join paths and filename: [dir1, dir2, filename.txt]")
	p := filepath.Join("dir1", "dir2", "filename.txt")
	pl(p)

	pl("# filepath will take care of separating symbols in path\n# we do some [dir//], or [dir1/../dir1]")
	pl(filepath.Join("dir1//", "filename"))
	pl(filepath.Join("dir1/../dir1", "filename"))

	pl("# after we have", p)
	pl("# try to get dir")
	pl("Dir(p):", filepath.Dir(p))
	pl("# get base name")
	pl("Base:", filepath.Base(p))
	pl("# get extension or suffix")
	pl("Ext:", filepath.Ext(p))

	pl("check some path if it's absolute")
	pl("# dir/file\n", filepath.IsAbs("dir/file"))
	pl("# /dir/file\n", filepath.IsAbs("/dir/file"))
	// not work so well on windows
	// pl("# C:\\dir\\file\n", filepath.IsAbs("C:\\dir\\path\n"))

	pl("# let's find relative path")
	pl("# 1st: a/b to a/b/t/file")
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	check(err)
	pl(rel)

	pl("# 2nd: a/b to a/c/t/file")
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	check(err)
	pl(rel)

	return nil
}

func (fp *FilePathExpl) Init() {
	if err := goexpl.RegisterGoExample(fp.GetGoExampleName(), &FilePathExpl{}); err != nil {
		panic(err)
	}
}

func (fp *FilePathExpl) GetGoExampleName() string {
	return common.FilePath
}
