package strexpl

import (
	"fmt"
	"strconv"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// go built-in string functions to parse string -> num or num -> string
type StrNumExpl struct{}

func (sne *StrNumExpl) RunExample(inputParams *goexpl.InputParams) error {
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("3.1415 -> %T: %f\n", f, f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Printf("123 -> %T: %d\n", i, i)

	// strconv can recognize hex num
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Printf("0x1c8 -> %T: %d\n", d, d)

	// if you want usigned int
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Printf("789 -> %T: %d\n", u, u)

	// Atoi is convenient if you just need 10-based int
	k, _ := strconv.Atoi("135")
	fmt.Printf("135 -> %T: %d\n", k, k)
	// err will be returned if you pass a invalid num string
	_, e := strconv.Atoi("wat")
	fmt.Println(e)

	return nil
}

func (sne *StrNumExpl) Init() {
	err := goexpl.RegisterGoExample(sne.GetGoExampleName(), &StrNumExpl{})
	if err != nil {
		panic(err)
	}
}

func (sne *StrNumExpl) GetGoExampleName() string {
	return common.StrNumExpl
}
