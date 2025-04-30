package timeexpl

import (
	"fmt"
	"time"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type TimeProcessExpl struct{}

func (tpe *TimeProcessExpl) RunExample(inputParams *goexpl.InputParams) error {
	p := fmt.Println

	t := time.Now()
	// basic example: use rfc3339 standard to format time
	p(t.Format(time.RFC3339))

	// use rfc3339 to parse time
	t1, err := time.Parse(
		time.RFC3339,
		"2035-10-01T08:00:00+08:00",
	)
	if err != nil {
		return err
	}
	p(t1)

	// use some custom layouts to display time
	// these string must be the time below
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:00 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, err := time.Parse(form, "5 44 PM")
	if err != nil {
		return err
	}
	p(t2)

	// use fmt to make your own output
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
	)

	// time.Parse return an error if time string is invalid
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)

	return nil
}

func (tpe *TimeProcessExpl) Init() {
	if err := goexpl.RegisterGoExample(tpe.GetGoExampleName(), &TimeProcessExpl{}); err != nil {
		panic(err)
	}
}

func (tpe *TimeProcessExpl) GetGoExampleName() string {
	return common.TimeProcessExpl
}
