package strexpl

import "fmt"

var (
	p  = fmt.Println
	pf = fmt.Printf
)

type Types struct {
	str      string
	hexStr   string
	shortStr string
	integer  int
	floatNum float32
	boolean  bool
}

var t Types = Types{
	boolean:  false,
	integer:  97,
	floatNum: 12340000.00,
	str:      "\"string\"",
	hexStr:   "hex this",
	shortStr: "foo",
}
