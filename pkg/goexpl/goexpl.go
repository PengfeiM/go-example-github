package goexpl

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type InputParams struct {
	argc int
	argv []string
	file os.File //预留参数，暂时未使用
}

func NewInputParams(argc int, argv []string, file os.File) *InputParams {
	return &InputParams{
		argc: argc,
		argv: argv,
		file: file,
	}
}

func (i *InputParams) GetArgNum() int {
	return i.argc
}

func (i *InputParams) GetArgs() []string {
	return i.argv
}

func (i *InputParams) GetFile() os.File {
	return i.file
}

type GoExample interface {
	RunExample(inputParams *InputParams) error
	Init()
	GetGoExampleName() string
}

var (
	goExampleMap      = map[string]GoExample{}
	goExampleNameList []string
)

func RegisterGoExample(name string, goExample GoExample) error {
	key := strings.ToLower(name)
	if _, ok := goExampleMap[key]; ok {
		return fmt.Errorf("Example %s has been regsitered", key)
	}
	goExampleMap[key] = goExample
	return nil
}

func GetGoExampleNameList() []string {
	return goExampleNameList
}

func GetGoExampleMap() map[string]GoExample {
	return goExampleMap
}

func Init() {

	// get map keys to make a list of go examples
	j := 0
	keys := make([]string, len(goExampleMap))
	for k := range goExampleMap {
		keys[j] = k
		j++
	}
	goExampleNameList = keys
	sort.Strings(goExampleNameList)
}
