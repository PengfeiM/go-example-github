package switchbranch

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"time"
)

type SwitchBranch struct {
}

func (s *SwitchBranch) GetGoExampleName() string {
	return common.SwitchBranch
}

func (s *SwitchBranch) RunExample(inputParams *goexpl.InputParams) error {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("null")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm an Integer")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	return nil
}

func (s *SwitchBranch) Init() {
	if err := goexpl.RegisterGoExample(s.GetGoExampleName(), &SwitchBranch{}); err != nil {
		panic(err)
	}
}
