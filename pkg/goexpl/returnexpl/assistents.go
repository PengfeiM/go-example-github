package returnexpl

import (
	"encoding/json"
	"fmt"
)

func func1() string {
	fmt.Println("this is func1")
	return "111"
}

func func2() string {
	fmt.Println("this is func2")
	return "222"
}

func returnFunc() (string, string) {
	fmt.Println("beginning of func")
	defer fmt.Println("end of func")
	return func2(), func1()
}

type output struct {
	Val1 int
	Val2 string
}

type executor struct {
	out interface{}
}

func newExecutor(data interface{}) *executor {
	exec := &executor{}
	exec.out = data
	return exec
}

var (
	out_json string = `{"Val1": 1, "Val2": "123"}`
)

func (e *executor) do() error {
	fmt.Println("executor doing")
	data := []byte(out_json)
	fmt.Println("data:", data)
	err := json.Unmarshal(data, &e.out)
	if err != nil {
		return err
	}
	fmt.Println("out:", e.out)
	return fmt.Errorf("nothing")
}

func modifyReturn() (*output, error) {
	out := &output{}
	exec := newExecutor(out)
	fmt.Println("initial value:\n\tout:", out, "\n\texec:", exec)
	defer func() {
		fmt.Println(out, exec)
	}()
	return out, exec.do()
}
