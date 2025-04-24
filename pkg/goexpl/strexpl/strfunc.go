package strexpl

import (
	s "strings"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type StrFuncExpl struct {
	a string
}

func (sf *StrFuncExpl) RunExample(inputParams *goexpl.InputParams) error {
	p("Contains:    test(s)             ->  ", s.Contains("test", "es"))
	p("Count:       test(t)             ->  ", s.Count("test", "t"))
	p("HasPrefix:   [te]st              ->  ", s.HasPrefix("test", "te"))
	p("HasSuffix:   te[st]              ->  ", s.HasSuffix("test", "st"))
	p("Index:       t[e]st              ->  ", s.Index("test", "e"))
	p("Join:        a[-]b               ->  ", s.Join([]string{"a", "b"}, "-"))
	p("Reqeat:      a * 5               ->  ", s.Repeat("a", 5))
	p("Replace:     foo(o->0)(-1)       ->  ", s.Replace("foo", "o", "0", -1))
	p("Replace:     foo(o->0)(1)        ->  ", s.Replace("foo", "o", "0", 1))
	p("ReplaceAll:  foo(o->0)(all)      ->  ", s.ReplaceAll("foo", "o", "0"))
	p("Split:       a[-]b[-]c[-]d[-]e   ->  ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:     TEST↓               ->  ", s.ToLower("TEST"))
	p("ToUpper:     test↑               ->  ", s.ToUpper("test"))

	return nil
}

func (sf *StrFuncExpl) Init() {
	if err := goexpl.RegisterGoExample(sf.GetGoExampleName(), &StrFuncExpl{}); err != nil {
		panic(err)
	}
}

func (sf *StrFuncExpl) GetGoExampleName() string {
	return common.StringFunc
}
