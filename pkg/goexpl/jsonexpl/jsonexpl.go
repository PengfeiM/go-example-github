package jsonexpl

import (
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// JsonExpl
// golang built-in json encoding and decoding tool
type JsonExpl struct{}

func (je *JsonExpl) RunExample(inputParams *goexpl.InputParams) error {
	// look into these two func, s'il vous plait
	encodeExample()
	decodeExample()

	return nil
}

func (je *JsonExpl) Init() {
	if err := goexpl.RegisterGoExample(je.GetGoExampleName(), &JsonExpl{}); err != nil {
		panic(err)
	}
}

func (je *JsonExpl) GetGoExampleName() string {
	return common.JsonExpl
}
