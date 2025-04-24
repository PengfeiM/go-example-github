package texttemp

import (
	"os"
	"text/template"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type TextTempExpl struct{}

func tt1() {
	// new a template instance, all you need to specified is its name
	t1 := template.New("t1")
	// use Parse() to make its body
	// {{.}} means a blank needs to be filled
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	// apart from using err!=nil to check template, you can also use .Must() to do it
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// use .Excute to pass a value to template
	// in this case, we output the result to stdout directly
	t1.Execute(os.Stdout, "some text for t1")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"c++",
		"c#",
	})
}

func tt2() {
	t2 := createTemp("t2", "Name: {{.Name}}\n")

	// use .Name to access exported field in struct
	t2.Execute(os.Stdout, struct {
		Name string
		Age  int
	}{"Jane Doe", 18})
	// works for map, too
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Michy Mouse",
	})
}

func (tt *TextTempExpl) RunExample(inputParams *goexpl.InputParams) error {
	// 1st text/template test
	tt1()
	tt2()
	return nil
}

func (tt *TextTempExpl) Init() {
	if err := goexpl.RegisterGoExample(tt.GetGoExampleName(), &TextTempExpl{}); err != nil {
		panic(err)
	}
}

func (tt *TextTempExpl) GetGoExampleName() string {
	return common.TextTemp
}
