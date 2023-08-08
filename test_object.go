package main

import (
	"fmt"
	"reflect"
)

var res = make(map[any]int)

func assertobject(objects ...interface{}) {
	fmt.Println(objects)
	for _, obj := range objects {
		fmt.Println(obj)
		if obj == nil {
			res["nil"]++
			continue
		}
		ref := reflect.TypeOf(obj).Kind()
		switch ref {
		case reflect.Array:
			res["Array"]++
		case reflect.Slice:
			res["Slice"]++
		case reflect.Map:
			res["Map"]++
		default:
			res[ref]++
		}
	}
}

func main() {
	var i int
	var c chan bool
	assertobject(uint(6), 5, "hi!", [3]int{3, 5, 8}, []int{1, 1}, []string{"bhb"},
		nil, 4.12, map[string]int{"j": 5, "k": 8}, &i, c, false)
	for k, v := range res {
		fmt.Printf("Total amount of %v type is %d\n", k, v)
	}
}
