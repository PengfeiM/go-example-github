package unittestexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type UnitTestExpl struct {
}

func (u *UnitTestExpl) RunExample(inputParams *goexpl.InputParams) error {
	// do nothing here
	fmt.Println("please run [go test ./pkg/goexpl/unittestexpl -run=Test_UpsertTestModelToDb " +
		"-coverprofile=coverage.out -gcflags=\"all=-N -l\"]\n and [go tool cover -html=coverage.out] to see result")
	return nil
}

func (u *UnitTestExpl) Init() {
	if err := goexpl.RegisterGoExample(u.GetGoExampleName(), &UnitTestExpl{}); err != nil {
		panic(err)
	}
}

func (u *UnitTestExpl) GetGoExampleName() string {
	return common.UnitTestExpl
}
