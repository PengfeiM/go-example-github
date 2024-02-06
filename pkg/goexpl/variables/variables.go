package variables

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type Variables struct {
}

func (v *Variables) GetGoExampleName() string {
	return common.Variables
}

func (v *Variables) RunExample(inputParams *goexpl.InputParams) error {
	// `var` declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	var e int
	fmt.Println(e)

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "apple"` in this case.
	// This syntax is only available inside functions.
	f := "apple"
	fmt.Println(f)

	return nil
}

func (v *Variables) Init() {
	err := goexpl.RegisterGoExample(v.GetGoExampleName(), &Variables{})
	if err != nil {
		return
	}
}
