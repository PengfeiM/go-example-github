package base64expl

import (
	b64 "encoding/base64"
	"fmt"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// Base64Expl
// go has built-in base64 process functions
type Base64Expl struct{}

func (b *Base64Expl) RunExample(inputParams *goexpl.InputParams) error {
	// init a string
	data := "abc123!?$*&()'-=@~"
	fmt.Println("original data:", data)

	// standard encoding
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("standard encoding:", sEnc)
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("standard decoding:", string(sDec))

	// url encoding
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("url-compatible encoding:", uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println("url-compatible decoding:", string(uDec))

	return nil
}

func (b *Base64Expl) Init() {
	if err := goexpl.RegisterGoExample(b.GetGoExampleName(), &Base64Expl{}); err != nil {
		panic(err)
	}
}

func (b *Base64Expl) GetGoExampleName() string {
	return common.Base64Expl
}
