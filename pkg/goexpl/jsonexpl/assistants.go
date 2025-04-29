package jsonexpl

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
   struct used in json encoding and decoding test, only Exported field can be manipulated by json
*/

// raw struct
type response1 struct {
	Page   int
	Fruits []string
}

// struct with json tags
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func encodeExample() {
	// first, some basic type
	// json.Marshal returns string in []byte
	// fmt.Println("json.Marshal(true)")
	fmt.Println("basic type encoded into json string([]byte)")
	bolB, _ := json.Marshal(true)
	fmt.Println("bool:\n", string(bolB)) // output it in string type
	intB, _ := json.Marshal(1)
	fmt.Println("int:\n", string(intB)) // output it in string type
	fltB, _ := json.Marshal(3.14)
	fmt.Println("float:\n", string(fltB)) // output it in string type
	strB, _ := json.Marshal("gopher")
	fmt.Println("string:\n", string(strB)) // output it in string type

	// slices and maps
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("slice:\n", string(slcB))
	mapD := map[string]int{
		"apple": 5,
		"peach": 9,
	}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("map:\n", string(mapB))

	// one biggest usage of json is to encode struct
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	fmt.Printf("struct: %+v\n", *res1D)
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// in golang, you can use tag to point json key names
	// look into response2's definition, s'il vous plait
	res2D := &response2{
		Page: 1,
		Fruits: []string{
			"apple",
			"peach",
			"banana",
		},
	}
	fmt.Printf("struct: %+v\n", *res2D)
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
}

func decodeExample() {
	fmt.Println("json decode test")
	// a byte slice for test
	byt := []byte(`{"num":6.13, "strs":["a","b"]}`)

	// a map
	var dat map[string]any

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]any)
	str1 := strs[0].(string)
	fmt.Println(str1)

	// decode json string to struct
	str := `{"page":1, "fruits":["apple","peach","pear"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// or you can use encoder and decoder in json
	// encode, which output result to stdout, you can change it to other io writer
	fmt.Println("encoder")
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// decoder
	fmt.Println("decoder")
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}
