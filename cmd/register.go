package cmd

import (
	"fmt"
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
	"go-example/pkg/goexpl/mapexpl"
	"go-example/pkg/goexpl/methodexpl"
	"go-example/pkg/goexpl/mutexexpl"
	"go-example/pkg/goexpl/panicexpl"
	"go-example/pkg/goexpl/pointer"
	"go-example/pkg/goexpl/pongo2expl"
	"go-example/pkg/goexpl/rangeexpl"
	"go-example/pkg/goexpl/ratelimitingexpl"
	"go-example/pkg/goexpl/recoverexpl"
	"go-example/pkg/goexpl/recursion"
	"go-example/pkg/goexpl/shiftexpl"
	"go-example/pkg/goexpl/slice"
	"go-example/pkg/goexpl/strrune"
	"go-example/pkg/goexpl/structexpl"
	"go-example/pkg/goexpl/switchbranch"
	"go-example/pkg/goexpl/timeexpl"
	"go-example/pkg/goexpl/values"
	"go-example/pkg/goexpl/variables"
	"go-example/pkg/goexpl/workerexpl"
	"reflect"
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
	MapExpl               *mapexpl.MapExpl
	MethodExpl            *methodexpl.MethodExpl
	MutexExpl             *mutexexpl.MutexExpl
	NonBlockSelectExpl    *channelexpl.NonBlockSelectExpl
	PanicExpl             *panicexpl.PanicExpl
	Pointer               *pointer.Pointer
	Pongo2Expl            *pongo2expl.Pongo2Expl
	RangeExpl             *rangeexpl.RangeExpl
	RateLimitingExpl      *ratelimitingexpl.RateLimitingExpl
	RecoverExpl           *recoverexpl.RecoverExpl
	Recursion             *recursion.Recursion
	SelectExpl            *channelexpl.SelectExpl
	ShiftExpl             *shiftexpl.ShiftExpl
	Slices                *slice.Slices
	SortingExpl           *slice.SortingExpl
	SortByFunc            *slice.SortByFuncExpl
	StatefulGoroutineExpl *goroutineexpl.StatefulGoroutineExpl
	StringRune            *strrune.StringRune
	StructEmbedding       *structexpl.Embedding
	StructExpl            *structexpl.StructExpl
	SwitchBranch          *switchbranch.SwitchBranch
	TimeOut               *channelexpl.TimeOutExpl
	TimerExpl             *timeexpl.TimerExpl
	TickerExpl            *timeexpl.TickerExpl
	Values                *values.Values
	Vars                  *variables.Variables
	WaitGroupExpl         *goroutineexpl.WaitGroupExpl
	WorkerExpl            *workerexpl.WorkerExpl
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
		MapExpl:               &mapexpl.MapExpl{},
		MethodExpl:            &methodexpl.MethodExpl{},
		MutexExpl:             &mutexexpl.MutexExpl{},
		NonBlockSelectExpl:    &channelexpl.NonBlockSelectExpl{},
		PanicExpl:             &panicexpl.PanicExpl{},
		Pointer:               &pointer.Pointer{},
		Pongo2Expl:            &pongo2expl.Pongo2Expl{},
		RangeExpl:             &rangeexpl.RangeExpl{},
		RateLimitingExpl:      &ratelimitingexpl.RateLimitingExpl{},
		RecoverExpl:           &recoverexpl.RecoverExpl{},
		Recursion:             &recursion.Recursion{},
		SelectExpl:            &channelexpl.SelectExpl{},
		ShiftExpl:             &shiftexpl.ShiftExpl{},
		Slices:                &slice.Slices{},
		SortByFunc:            &slice.SortByFuncExpl{},
		SortingExpl:           &slice.SortingExpl{},
		StatefulGoroutineExpl: &goroutineexpl.StatefulGoroutineExpl{},
		StringRune:            &strrune.StringRune{},
		StructEmbedding:       &structexpl.Embedding{},
		StructExpl:            &structexpl.StructExpl{},
		SwitchBranch:          &switchbranch.SwitchBranch{},
		TimeOut:               &channelexpl.TimeOutExpl{},
		TimerExpl:             &timeexpl.TimerExpl{},
		TickerExpl:            &timeexpl.TickerExpl{},
		Values:                &values.Values{},
		Vars:                  &variables.Variables{},
		WaitGroupExpl:         &goroutineexpl.WaitGroupExpl{},
		WorkerExpl:            &workerexpl.WorkerExpl{},
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
