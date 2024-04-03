package cmd

import (
	"fmt"
	"go-example/pkg/goexpl/arrays"
	"go-example/pkg/goexpl/channelexpl"
	"go-example/pkg/goexpl/closure"
	"go-example/pkg/goexpl/constants"
	"go-example/pkg/goexpl/cronjob"
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
	"go-example/pkg/goexpl/pointer"
	"go-example/pkg/goexpl/pongo2expl"
	"go-example/pkg/goexpl/rangeexpl"
	"go-example/pkg/goexpl/recursion"
	"go-example/pkg/goexpl/slice"
	"go-example/pkg/goexpl/strrune"
	"go-example/pkg/goexpl/structexpl"
	"go-example/pkg/goexpl/switchbranch"
	"go-example/pkg/goexpl/values"
	"go-example/pkg/goexpl/variables"
	"reflect"
)

type AllExampleList struct {
	Arrays               *arrays.Arrays
	ChannelBufferExpl    *channelexpl.ChannelBufferExpl
	ChannelCloseExpl     *channelexpl.ChannelCloseExpl
	ChannelDirectionExpl *channelexpl.ChannelDirectionExpl
	ChannelExpl          *channelexpl.ChannelExpl
	ChannelRangeExpl     *channelexpl.ChannelRangeExpl
	ChannelSyncExpl      *channelexpl.ChannelSyncExpl
	Constants            *constants.Constants
	Closure              *closure.Closure
	CronJob              *cronjob.CronJob
	ErrorExpl            *errorexpl.ErrorExpl
	ForExpl              *forexpl.ForExpl
	FuncExpl             *funcexpl.FuncExpl
	Generics             *generics.Generics
	GoRoutineExpl        *goroutineexpl.GoRoutineExpl
	HelloWorld           *helloworld.HelloWorld
	IfEl                 *ifel.IfEl
	InterfaceExpl        *interfaceexpl.InterfaceExpl
	MapExpl              *mapexpl.MapExpl
	MethodExpl           *methodexpl.MethodExpl
	NonBlockSelectExpl   *channelexpl.NonBlockSelectExpl
	Pointer              *pointer.Pointer
	Pongo2Expl           *pongo2expl.Pongo2Expl
	RangeExpl            *rangeexpl.RangeExpl
	Recursion            *recursion.Recursion
	SelectExpl           *channelexpl.SelectExpl
	Slices               *slice.Slices
	StringRune           *strrune.StringRune
	StructEmbedding      *structexpl.Embedding
	StructExpl           *structexpl.StructExpl
	SwitchBranch         *switchbranch.SwitchBranch
	TimeOut              *channelexpl.TimeOutExpl
	Values               *values.Values
	Vars                 *variables.Variables
}

func allExampleList() *AllExampleList {
	return &AllExampleList{
		Arrays:               &arrays.Arrays{},
		ChannelBufferExpl:    &channelexpl.ChannelBufferExpl{},
		ChannelCloseExpl:     &channelexpl.ChannelCloseExpl{},
		ChannelDirectionExpl: &channelexpl.ChannelDirectionExpl{},
		ChannelExpl:          &channelexpl.ChannelExpl{},
		ChannelRangeExpl:     &channelexpl.ChannelRangeExpl{},
		ChannelSyncExpl:      &channelexpl.ChannelSyncExpl{},
		Constants:            &constants.Constants{},
		Closure:              &closure.Closure{},
		CronJob:              &cronjob.CronJob{},
		ErrorExpl:            &errorexpl.ErrorExpl{},
		ForExpl:              &forexpl.ForExpl{},
		FuncExpl:             &funcexpl.FuncExpl{},
		Generics:             &generics.Generics{},
		GoRoutineExpl:        &goroutineexpl.GoRoutineExpl{},
		HelloWorld:           &helloworld.HelloWorld{},
		IfEl:                 &ifel.IfEl{},
		InterfaceExpl:        &interfaceexpl.InterfaceExpl{},
		MapExpl:              &mapexpl.MapExpl{},
		MethodExpl:           &methodexpl.MethodExpl{},
		NonBlockSelectExpl:   &channelexpl.NonBlockSelectExpl{},
		Pointer:              &pointer.Pointer{},
		Pongo2Expl:           &pongo2expl.Pongo2Expl{},
		RangeExpl:            &rangeexpl.RangeExpl{},
		Recursion:            &recursion.Recursion{},
		SelectExpl:           &channelexpl.SelectExpl{},
		Slices:               &slice.Slices{},
		StringRune:           &strrune.StringRune{},
		StructEmbedding:      &structexpl.Embedding{},
		StructExpl:           &structexpl.StructExpl{},
		SwitchBranch:         &switchbranch.SwitchBranch{},
		TimeOut:              &channelexpl.TimeOutExpl{},
		Values:               &values.Values{},
		Vars:                 &variables.Variables{},
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
