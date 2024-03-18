package pongo2expl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"

	"github.com/flosch/pongo2"
)

type Pongo2Expl struct {
	templateFile string
	metaDataFile string
	outputFile   string
}

func (pge *Pongo2Expl) RunExample(inputParams *goexpl.InputParams) error {

	_, err := pongo2.FromFile(pge.templateFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("not implemented")

	return nil
}

func (pge *Pongo2Expl) Init() {
	if err := goexpl.RegisterGoExample(pge.GetGoExampleName(), &Pongo2Expl{
		templateFile: "./files/pongo2.tpl",
		metaDataFile: "./files/pongo2.json",
		outputFile:   "./files/pongo2.output",
	}); err != nil {
		panic(err)
	}
}

func (pge *Pongo2Expl) GetGoExampleName() string {
	return common.Pongo2Expl
}
