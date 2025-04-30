package randexpl

import (
	"fmt"
	"math/rand/v2"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// go have random package as math/rand/v2, which is newer than math/rand
type RandExpl struct{}

func (r *RandExpl) RunExample(inputParams *goexpl.InputParams) error {
	// get a rand integer
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// get a rand float64 in [0.0,1.0)
	fmt.Println(rand.Float64())
	// do some simple math, if you wish to get other random float
	fmt.Println((rand.Float64()*5 + 5), ",", (rand.Float64()*5 + 5))

	// generate your own seed
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Println(r2.IntN(100), ",", r2.IntN(100))
	// if you pass the same seed, you will get the very same rand sequence
	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Println(r3.IntN(100), ",", r3.IntN(100))

	return nil
}

func (r *RandExpl) Init() {
	if err := goexpl.RegisterGoExample(r.GetGoExampleName(), &RandExpl{}); err != nil {
		panic(err)
	}
}

func (r *RandExpl) GetGoExampleName() string {
	return common.RandExpl
}
