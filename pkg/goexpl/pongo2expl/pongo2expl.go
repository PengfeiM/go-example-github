package pongo2expl

import (
	"encoding/json"
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"os"

	"github.com/flosch/pongo2"
)

type Pongo2Expl struct {
	templateFile string
	metaDataFile string
	outputFile   string
}

//type metaData struct {
//Name       string `json:"name"`
//Occupation string `json:"occupation"`
//ABool      bool   `json:"aBool"`
//}

func (pge *Pongo2Expl) RunExample(inputParams *goexpl.InputParams) error {

	tpl, err := pongo2.FromFile(pge.templateFile)
	if err != nil {
		panic(err)
	}

	paramFile, err := os.Open(pge.metaDataFile)
	if err != nil {
		panic("open metaDataFile failed")
	}
	defer paramFile.Close()
	jsonDecoder := json.NewDecoder(paramFile)

	//var md metaData

	//err = jsonDecoder.Decode(&md)
	//if err != nil {
	//return err
	//}

	//fmt.Println(md)

	//ctx := pongo2.Context{"metas": md}

	var data map[string]interface{}
	jsonDecoder.Decode(&data)
	fmt.Println(data)

	ctx := pongo2.Context(data)
	res, err := tpl.Execute(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("渲染结果")
	fmt.Println(res)

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
