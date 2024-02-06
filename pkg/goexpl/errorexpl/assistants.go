package errorexpl

import (
	"errors"
	"fmt"
)

// f1
// 习惯上来说，一般使用函数最后一个返回值返回错误，类型为 error，built-in 的一个接口
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New() 返回一个基础的 error
		return -1, errors.New("can't work with 42")
	}
	// use nil value to show there is no error
	return arg + 3, nil
}

// argError
// implement interface error as argError
//	type error interface {
//		Error() string
//	}
// By this way, user can define their own type of error
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// use our own error type to paas two field into error
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
