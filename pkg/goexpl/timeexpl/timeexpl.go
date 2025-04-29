package timeexpl

import (
	"fmt"
	"time"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// TimeExpl
// demonstrate go's built-in time functions
type TimeExpl struct{}

func (t *TimeExpl) RunExample(inputParams *goexpl.InputParams) error {
	p := fmt.Println

	// get current time: 2025-04-29 19:34:48.373149 +0800 CST m=+0.002313251
	now := time.Now()
	p(now)

	// build a time with year, month, day, etc.
	// and do not forget a timezone.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 123456789, time.UTC,
	)
	// 2009-11-17 20:34:58.123456789 +0000 UTC
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	// which weekday is it?
	p(then.Weekday())

	// how to compare time?
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// calculate gap between time
	diff := now.Sub(then)
	// 135399h7m31.501789211s
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// add & minus for time
	// minus is add -time
	p(then.Add(diff))
	p(then.Add(-diff))

	return nil
}

func (t *TimeExpl) Init() {
	if err := goexpl.RegisterGoExample(t.GetGoExampleName(), &TimeExpl{}); err != nil {
		panic(err)
	}
}

func (t *TimeExpl) GetGoExampleName() string {
	return common.TimeExpl
}
