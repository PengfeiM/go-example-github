package generics

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// golang 从 1.18 版本开始支持泛型，即 type 参数
type Generics struct {
}

func (g *Generics) RunExample(inputParams *goexpl.InputParams) error {
	var m = map[int]string{1: "2", 2: "4", 3: "8"}
	// 一般来说，使用泛型声明的函数，我们在调用时，编译器会自动为我们确认入参的类型。
	fmt.Println(MapKeys(m))

	// 当然，也可以手动指定类型
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list: ", lst.GetAll())

	return nil
}

func (g *Generics) Init() {
	err := goexpl.RegisterGoExample(g.GetGoExampleName(), &Generics{})
	if err != nil {
		panic(err)
	}
}

func (g *Generics) GetGoExampleName() string {
	return common.Generics
}

// MapKeys
/*
	这是一个泛型函数的示例。它接收指定类型范围的 map 并输出 key 的 slice。
	这个函数有两个 type 参数，K 和 V，K 被限制为 comparable 类型，即可以进行 ==
	和 != 比较。这也是 go 中 map 的 key 所需的类型限制。V 为 any 类型，也就是可以为任意类型
	comparable 的定义：
		type comparable interface{ comparable }
	any 的定义：
		type any = interface{}
*/
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// 泛型的另外一个应用，可以用于声明一个元素为任意类型的单链表
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Push
// 为单链表 List[T] 定义的方法中，也要声明类型参数
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}
