// Package slice
/*
slice 是 golang 语言中一个重要的序列化数据类型
*/
package slice

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type Slices struct {
}

func (slc *Slices) RunExample(inputParams *goexpl.InputParams) error {
	/*
		不像 array，长度并不是它的 type 属性，只有元素类型才是。
		未初始化的 slice 的值为 nil，且长度为 0
	*/
	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s) == 0)

	/*
		为了创建一个长度不为 0 的 slice，我们可以使用内置的 make 函数。
		比如创建一个长度为 3 的 slice。
		其元素默认值也是零值。
		如果在设计程序时知道 slice 可能的极限长度，那么为 slice 预分配一个长度是有益的。
	*/
	s = make([]string, 3)
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s))

	/*
		slice 具有与 array 类似的操作方式
	*/
	s[0] = "a"
	s[1] = "ab"
	s[2] = "abc"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	// len:  3 cap:  3
	fmt.Println("len: ", len(s), "cap: ", cap(s))

	/*
		slice 还具备一些 array 所不具备的特性，比如可以用内置的 append() 函数来在 slice 尾部增加新元素
		从下面的🌰可以看出，slice 的长度随着 append 的进行，逐渐增长。
		但是需要注意的事，slice 的容量（cap）并不是和 length 一样，线性增长的。
		而是呈指数增长，每次 cap 不够用时，new cap = old cap * 2。
	*/
	s = append(s, "abcd") // len:  4 cap:  6
	fmt.Println("len: ", len(s), "cap: ", cap(s))
	s = append(s, "abcde", "abcdef") // len:  6 cap:  6
	fmt.Println("len: ", len(s), "cap: ", cap(s))
	s = append(s, "abcdefg") // len:  7 cap:  12
	fmt.Println("apd: ", s)
	fmt.Println("len: ", len(s), "cap: ", cap(s))

	/*
		slice 还支持 copy 操作
		这里你会问了，为什么要用 copy，而不是直接使用“=”给一个新的 slice 赋值呢？
		这里涉及 slice 的浅复制和深复制
		type slice struct{
			point unsafe.Pointer
			len int
			cap int
		}
	*/
	c1 := s                      // shallow copy
	c2 := make([]string, len(s)) // 为 c 分配一个和 s 一样大的空间
	copy(c2, s)                  // copy s 中的值到 c 中, deep copy
	fmt.Printf("s: %v\nc1: %v\nc2: %v\n", s, c1, c2)
	fmt.Printf("s addr: %p\nc1 addr: %p\nc2 addr: %p\n", &s, &c1, &c2) // 可以看到 s、c1、c2的地址（slice 结构体地址）并不相同
	/*
		记录一次 debug 记录
		地址：
			s: 0x1400000e3c0-0x1400000e3d7
			c1: 0x14000182018-0x1400018202f
			c2: 0x14000182030-0x14000182047
		内容：
				ptr				len			cap
		s		0x14000180000	0x07		0x0c
		c1		0x14000180000	0x07		0x0c
		c2		0x14000186000	0x07		0x07
		对比可以发现，c1 中的 *string 类型指针和 s 中的 *string 类型指针所存储的地址是相同的；c2 则不同

		检查 s 的内容
		[]string 类型中，ptr 指向一个 string 类型的结构体
		type string struct{
			str unsafe.Pointer
			len int
		}
		比如上面的 s 中， ptr 指向了一个 string 结构体
				str-ptr			len
		ptr		0x0104928699	0x01
		str-ptr 这个指针指向真正的字符串（ascii/utf-8存储的字符）所在地址空间的首地址，检查其内容可以发现是：
		0x61 = a，符合预期
		同理，c2 中的 *string 指针0x14000186000 指向了一个 *byte 指针 0x0104928699
		虽然，*string 指针内容不同，但是因为字符串内容相同，真正的指向 *byte （存储ascii码）的指针居然又相同了。
		奇怪的优化，如果这里不是 string 的 slice 情况应该就不是这样了。
	*/
	// 在改变 s[0] 之前，让我们再测试一个东西：相等？
	// 鉴于禁止 slice 进行 == 操作，只能用 slices.Equal()函数了
	s[0] = "1234567"                                 // modify s[0]
	fmt.Printf("s: %v\nc1: %v\nc2: %v\n", s, c1, c2) // s 和 c1 中的指针指向地址相同。len、cap也一样。c2则不然
	fmt.Println("s cap: ", cap(s), "c1 cap: ", cap(c1), "c2 cap: ", cap(c2))
	fmt.Println("s len: ", len(s), "c1 len: ", len(c1), "c2 len: ", cap(c2))

	/*
		slice 作为“切片”，支持“子切片”
	*/
	l := s[2:5] // index1 -> index2
	fmt.Println("sl1: ", l)

	l = s[:5] // head -> index2
	fmt.Println("sl2: ", l)

	l = s[2:] // index1 -> tail
	fmt.Println("sl3: ", l)

	// slice 的初始化也可以类似array
	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	// 不像 array，2d 的slice，内层 slice 的长度是可以不一致的。
	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	return nil
}

func (slc *Slices) Init() {
	if err := goexpl.RegisterGoExample(slc.GetGoExampleName(), &Slices{}); err != nil {
		panic(err)
	}
}

func (slc *Slices) GetGoExampleName() string {
	return common.Slices
}
