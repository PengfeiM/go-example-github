package values

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type Values struct {
}

func (val *Values) GetGoExampleName() string {
	return common.Values
}

func (val *Values) RunExample(inputParams *goexpl.InputParams) error {
	fmt.Println("# Run example values...")
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println("# Example Over! ")
	return nil
}

func (val *Values) Init() {
	if err := goexpl.RegisterGoExample(val.GetGoExampleName(), &Values{}); err != nil {
		panic(err)
	}
}
