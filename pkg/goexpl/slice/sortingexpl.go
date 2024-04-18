package slice

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"golang.org/x/exp/slices"
)

// SortingExpl
// Go's slices package provides some sorting methods
type SortingExpl struct {
}

func (se *SortingExpl) RunExample(inputParams *goexpl.InputParams) error {
	strs := []string{"123", "abc", "bac", "bca"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{123, 555, 321, 222}

	slices.Sort(ints)
	fmt.Println("Integers:", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted?", s)

	return nil
}

func (se *SortingExpl) Init() {
	if err := goexpl.RegisterGoExample(se.GetGoExampleName(), &SortingExpl{}); err != nil {
		panic(err)
	}
}

func (se *SortingExpl) GetGoExampleName() string {
	return common.SortingExpl
}
