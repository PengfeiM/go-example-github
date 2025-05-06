package shaexpl

import (
	"crypto/sha256"
	"fmt"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// SHAExpl goexpl.shaexpl.SHAExpl
// sha256 hash is a frequently used to compute short identities for binary or text blobs.
// golang has a built-in sha256 support
type SHAExpl struct{}

func (se *SHAExpl) RunExample(inputParams *goexpl.InputParams) error {
	// a string
	s := "sha256 this string"
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	return nil
}

func (se *SHAExpl) Init() {
	if err := goexpl.RegisterGoExample(se.GetGoExampleName(), &SHAExpl{}); err != nil {
		panic(err)
	}
}

func (se *SHAExpl) GetGoExampleName() string {
	return common.SHAExpl
}
