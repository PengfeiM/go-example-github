package strexpl

import (
	"fmt"
	"os"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type StrFormatExpl struct{}

func (s *StrFormatExpl) RunExample(inputParams *goexpl.InputParams) error {
	// print struct
	pf("struct  ->  raw:    %%v     %v\n", t)
	pf("struct  ->  name:   %%+v    %+v\n", t)
	pf("struct  ->  syntax: %%#v    %#v\n", t)
	// print integer
	pf("int     ->  int:    %%d     %d\n", t.integer)
	pf("int     ->  binary: %%b     %b\n", t.integer)
	// num -> char
	pf("int     ->  char:   %%c     %c\n", t.integer)
	// hex
	pf("int     ->  hex:    %%x     %x\n", t.integer)
	// float
	pf("float   ->  float:  %%f     %f\n", t.floatNum)
	pf("float   ->  sciform:%%e     %e\n", t.floatNum)
	pf("float   ->  sciform:%%E     %E\n", t.floatNum)
	// string
	pf("string  ->  basic:  %%s     %s\n", t.str)
	pf("string  ->  quote:  %%q     %q\n", t.str)
	pf("string  ->  hex:    %%x     %x\n", t.str)
	// pointer
	pf("pointer ->  pointer:%%p     %p\n", &t)

	// advanced format
	// align
	pf("align1  ->  integer:        |%6d|%6d|\n", t.integer, t.integer*100)
	pf("align2  ->  float:          |%6.2f|%6.2f|\n", t.floatNum, t.floatNum/2333333)
	pf("align3  ->  left:           |%-6.2f|%-6.2f|\n", t.floatNum, t.floatNum/2333333)
	pf("align4  ->  string:         |%6s|%6s|\n", t.shortStr, t.shortStr)
	pf("align5  ->  str-left:       |%-6s|%-6s|\n", t.shortStr, t.shortStr)
	// special format
	str := fmt.Sprintf("sprintf: a %s", "string")
	p(str)
	fmt.Fprintf(os.Stderr, "io: as %s\n", "error")

	return nil
}

func (s *StrFormatExpl) Init() {
	err := goexpl.RegisterGoExample(s.GetGoExampleName(), &StrFormatExpl{})
	if err != nil {
		panic(err)
	}
}

func (s *StrFormatExpl) GetGoExampleName() string {
	return common.StringFormat
}
