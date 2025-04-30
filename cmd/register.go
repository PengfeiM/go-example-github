package cmd

import (
	"fmt"
	"reflect"

	"go-example/pkg/goexpl/arrays"
	"go-example/pkg/goexpl/atomiccounterexpl"
	"go-example/pkg/goexpl/channelexpl"
	"go-example/pkg/goexpl/closure"
	"go-example/pkg/goexpl/constantsexpl"
	"go-example/pkg/goexpl/cronjob"
	"go-example/pkg/goexpl/deferexpl"
	"go-example/pkg/goexpl/errorexpl"
	"go-example/pkg/goexpl/forexpl"
	"go-example/pkg/goexpl/funcexpl"
	"go-example/pkg/goexpl/generics"
	"go-example/pkg/goexpl/goroutineexpl"
	"go-example/pkg/goexpl/helloworld"
	"go-example/pkg/goexpl/ifel"
	"go-example/pkg/goexpl/interfaceexpl"
	"go-example/pkg/goexpl/jsonexpl"
	"go-example/pkg/goexpl/mapexpl"
	"go-example/pkg/goexpl/methodexpl"
	"go-example/pkg/goexpl/mutexexpl"
	"go-example/pkg/goexpl/netexpl"
	"go-example/pkg/goexpl/panicexpl"
	"go-example/pkg/goexpl/pointer"
	"go-example/pkg/goexpl/pongo2expl"
	"go-example/pkg/goexpl/randexpl"
	"go-example/pkg/goexpl/rangeexpl"
	"go-example/pkg/goexpl/ratelimitingexpl"
	"go-example/pkg/goexpl/recoverexpl"
	"go-example/pkg/goexpl/recursion"
	"go-example/pkg/goexpl/regexpexpl"
	"go-example/pkg/goexpl/returnexpl"
	"go-example/pkg/goexpl/shiftexpl"
	"go-example/pkg/goexpl/slice"
	"go-example/pkg/goexpl/strexpl"
	"go-example/pkg/goexpl/strrune"
	"go-example/pkg/goexpl/structexpl"
	"go-example/pkg/goexpl/switchbranch"
	"go-example/pkg/goexpl/texttemp"
	"go-example/pkg/goexpl/timeexpl"
	"go-example/pkg/goexpl/unittestexpl"
	"go-example/pkg/goexpl/values"
	"go-example/pkg/goexpl/variables"
	"go-example/pkg/goexpl/workerexpl"
)

type AllExampleList struct {
	Arrays                *arrays.Arrays
	AtomicCounterExpl     *atomiccounterexpl.AtomicCounterExpl
	ChannelBufferExpl     *channelexpl.ChannelBufferExpl
	ChannelCloseExpl      *channelexpl.ChannelCloseExpl
	ChannelDirectionExpl  *channelexpl.ChannelDirectionExpl
	ChannelExpl           *channelexpl.ChannelExpl
	ChannelRangeExpl      *channelexpl.ChannelRangeExpl
	ChannelSyncExpl       *channelexpl.ChannelSyncExpl
	Constants             *constantsexpl.Constants
	Closure               *closure.Closure
	CronJob               *cronjob.CronJob
	DeferExpl             *deferexpl.DeferExpl
	ErrorExpl             *errorexpl.ErrorExpl
	ForExpl               *forexpl.ForExpl
	FuncExpl              *funcexpl.FuncExpl
	Generics              *generics.Generics
	GoRoutineExpl         *goroutineexpl.GoRoutineExpl
	HelloWorld            *helloworld.HelloWorld
	IfEl                  *ifel.IfEl
	InterfaceExpl         *interfaceexpl.InterfaceExpl
	IpAddrExpl            *netexpl.IpAddrExpl
	JsonExpl              *jsonexpl.JsonExpl
	MapExpl               *mapexpl.MapExpl
	MethodExpl            *methodexpl.MethodExpl
	MutexExpl             *mutexexpl.MutexExpl
	NetExpl               *netexpl.NetExpl
	NonBlockSelectExpl    *channelexpl.NonBlockSelectExpl
	PanicExpl             *panicexpl.PanicExpl
	Pointer               *pointer.Pointer
	Pongo2Expl            *pongo2expl.Pongo2Expl
	RangeExpl             *rangeexpl.RangeExpl
	RateLimitingExpl      *ratelimitingexpl.RateLimitingExpl
	RecoverExpl           *recoverexpl.RecoverExpl
	Recursion             *recursion.Recursion
	RegExpExpl            *regexpexpl.RegExpExpl
	ReturnExpl            *returnexpl.ReturnExpl
	SelectExpl            *channelexpl.SelectExpl
	ShiftExpl             *shiftexpl.ShiftExpl
	Slices                *slice.Slices
	SortingExpl           *slice.SortingExpl
	SortByFunc            *slice.SortByFuncExpl
	StatefulGoroutineExpl *goroutineexpl.StatefulGoroutineExpl
	StringFormat          *strexpl.StrFormatExpl
	StringFunc            *strexpl.StrFuncExpl
	StringRune            *strrune.StringRune
	StructEmbedding       *structexpl.Embedding
	StructExpl            *structexpl.StructExpl
	SwitchBranch          *switchbranch.SwitchBranch
	TextTemp              *texttemp.TextTempExpl
	TimeOut               *channelexpl.TimeOutExpl
	TimerExpl             *timeexpl.TimerExpl
	TickerExpl            *timeexpl.TickerExpl
	UnitTestExpl          *unittestexpl.UnitTestExpl
	Values                *values.Values
	Vars                  *variables.Variables
	WaitGroupExpl         *goroutineexpl.WaitGroupExpl
	WorkerExpl            *workerexpl.WorkerExpl
	TimeExpl              *timeexpl.TimeExpl
	TimeProcessExpl       *timeexpl.TimeProcessExpl
	RandExpl              *randexpl.RandExpl
	StrNumExpl            *strexpl.StrNumExpl
}

func allExampleList() *AllExampleList {
	return &AllExampleList{
		Arrays:                &arrays.Arrays{},
		AtomicCounterExpl:     &atomiccounterexpl.AtomicCounterExpl{},
		ChannelBufferExpl:     &channelexpl.ChannelBufferExpl{},
		ChannelCloseExpl:      &channelexpl.ChannelCloseExpl{},
		ChannelDirectionExpl:  &channelexpl.ChannelDirectionExpl{},
		ChannelExpl:           &channelexpl.ChannelExpl{},
		ChannelRangeExpl:      &channelexpl.ChannelRangeExpl{},
		ChannelSyncExpl:       &channelexpl.ChannelSyncExpl{},
		Constants:             &constantsexpl.Constants{},
		Closure:               &closure.Closure{},
		CronJob:               &cronjob.CronJob{},
		DeferExpl:             &deferexpl.DeferExpl{},
		ErrorExpl:             &errorexpl.ErrorExpl{},
		ForExpl:               &forexpl.ForExpl{},
		FuncExpl:              &funcexpl.FuncExpl{},
		Generics:              &generics.Generics{},
		GoRoutineExpl:         &goroutineexpl.GoRoutineExpl{},
		HelloWorld:            &helloworld.HelloWorld{},
		IfEl:                  &ifel.IfEl{},
		InterfaceExpl:         &interfaceexpl.InterfaceExpl{},
		IpAddrExpl:            &netexpl.IpAddrExpl{},
		JsonExpl:              &jsonexpl.JsonExpl{},
		MapExpl:               &mapexpl.MapExpl{},
		MethodExpl:            &methodexpl.MethodExpl{},
		MutexExpl:             &mutexexpl.MutexExpl{},
		NetExpl:               &netexpl.NetExpl{},
		NonBlockSelectExpl:    &channelexpl.NonBlockSelectExpl{},
		PanicExpl:             &panicexpl.PanicExpl{},
		Pointer:               &pointer.Pointer{},
		Pongo2Expl:            &pongo2expl.Pongo2Expl{},
		RangeExpl:             &rangeexpl.RangeExpl{},
		RateLimitingExpl:      &ratelimitingexpl.RateLimitingExpl{},
		RecoverExpl:           &recoverexpl.RecoverExpl{},
		Recursion:             &recursion.Recursion{},
		RegExpExpl:            &regexpexpl.RegExpExpl{},
		ReturnExpl:            &returnexpl.ReturnExpl{},
		SelectExpl:            &channelexpl.SelectExpl{},
		ShiftExpl:             &shiftexpl.ShiftExpl{},
		Slices:                &slice.Slices{},
		SortByFunc:            &slice.SortByFuncExpl{},
		SortingExpl:           &slice.SortingExpl{},
		StatefulGoroutineExpl: &goroutineexpl.StatefulGoroutineExpl{},
		StringFormat:          &strexpl.StrFormatExpl{},
		StringFunc:            &strexpl.StrFuncExpl{},
		StringRune:            &strrune.StringRune{},
		StructEmbedding:       &structexpl.Embedding{},
		StructExpl:            &structexpl.StructExpl{},
		SwitchBranch:          &switchbranch.SwitchBranch{},
		TextTemp:              &texttemp.TextTempExpl{},
		TimeOut:               &channelexpl.TimeOutExpl{},
		TimerExpl:             &timeexpl.TimerExpl{},
		TickerExpl:            &timeexpl.TickerExpl{},
		UnitTestExpl:          &unittestexpl.UnitTestExpl{},
		Values:                &values.Values{},
		Vars:                  &variables.Variables{},
		WaitGroupExpl:         &goroutineexpl.WaitGroupExpl{},
		WorkerExpl:            &workerexpl.WorkerExpl{},
		TimeExpl:              &timeexpl.TimeExpl{},
		TimeProcessExpl:       &timeexpl.TimeProcessExpl{},
		RandExpl:              &randexpl.RandExpl{},
		StrNumExpl:            &strexpl.StrNumExpl{},
	}
}

func RegisterAll() {
	/*
		注册 Examples
	*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from Err:\n", r)
		}
	}()

	exampleList := allExampleList()
	t := reflect.TypeOf(exampleList)
	v := reflect.ValueOf(exampleList)

	for i := 0; i < t.Elem().NumField(); i++ {
		expl := t.Elem().Field(i)
		explValue := v.Elem().FieldByName(expl.Name)

		initFunc := explValue.MethodByName("Init")
		initFunc.Call([]reflect.Value{})
	}
}
