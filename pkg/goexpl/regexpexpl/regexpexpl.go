package regexpexpl

import (
	"bytes"
	"regexp"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// RegExpExpl: go has a built-in regular expression. use it by import "regexp"
type RegExpExpl struct{}

func (r *RegExpExpl) RunExample(inputParams *goexpl.InputParams) error {
	// use regexp to match string directly
	/*
		[a-z]: lower letter from "a" to "z"
		[a-z]+: 1 or multi lower letter
		xxx()xxx: () is a submatch
	*/
	match, _ := regexp.MatchString("p[a-z]+ch", "peach")
	p("reg: p[a-z]+ch contains peach?")
	p(match)

	// normally, a regexp could be used in multi times
	// so we use Compile to create a optimized reg struct
	reg, _ := regexp.Compile("p([a-z]+)ch")
	pf("here is a regexp instance we create: \n\t%+v\n", reg)

	p("reg matches \"peach\"?")
	p(reg.MatchString("peach"))

	p("can we find a match in string \"peach punch\"?")
	p(reg.FindString("peach punch"))
	p("index:", reg.FindStringIndex("peach punch")) // return match substring's start index and end index+1
	// p("peach"[0], "peach punch"[4])
	// find match p[a-z]+ch and submatch [a-z]+
	p("find whole match and submatch:", reg.FindStringSubmatch("peach punch"))
	p("find submatch's index:", reg.FindStringSubmatchIndex("peach punch"))

	// we can also try to find all matches in string
	// 									string				 num -> how many you need?
	p("all matches:", reg.FindAllString("peach punch pinch", -1))
	p("all matches:", reg.FindAllString("peach punch pinch", 2))

	// find all submatch, I'm sure that you have already know all the params' meaning
	p("all:", reg.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// by use reg.Match, we need to input a param in []byte
	p("match pattern in []byte type: []byte(\"peach\")")
	p(reg.Match([]byte("peach")))

	// different from Compile, MustCompile panic instead of returning an error
	// it's suitable when we use a regexp as global variable
	reg = regexp.MustCompile("p([a-z]+)ch")
	p("regexp:", reg)

	// apart from find, we can also replace with reg
	p("replace match with another string:", reg.ReplaceAllString("a peach", "<fruit>"))
	// point a func to do replace
	in := []byte("a peach")
	out := reg.ReplaceAllFunc(in, bytes.ToUpper)
	p("replace with bytes.ToUpper for string \"a peach\"")
	p(string(out))

	return nil
}

func (r *RegExpExpl) Init() {
	if err := goexpl.RegisterGoExample(r.GetGoExampleName(), &RegExpExpl{}); err != nil {
		panic(err)
	}
}

func (r *RegExpExpl) GetGoExampleName() string {
	return common.RegExpExpl
}
