package main

import (
	"fmt"
	"reflect"
)

var (
	totalNil int
	totalInt int
	totalFlt int
	totalStr int
	totalMap int
	totalSlc int
	totalArr int
	totalPtr int
	totalBln int
	totalRue int
	totalOtr int
)

func doCheck(i interface{}) {
	switch j := i.(type) {
	case nil:
		totalNil++
	case int:
		totalInt++
	case float64:
		totalFlt++
	case string:
		totalStr++
	case uintptr:
		totalPtr++
	case bool:
		totalBln++
	default:
		if reflect.TypeOf(j).Kind() == reflect.Slice {
			totalSlc++
			break
		}
		if reflect.TypeOf(j).Kind() == reflect.Array {
			totalArr++
			break
		}
		if reflect.TypeOf(j).Kind() == reflect.Map {
			totalMap++
			break
		}
		totalOtr++
	}
}

func assertobject(objects ...interface{}) {
	for _, obj := range objects {
		doCheck(obj)
	}
	fmt.Println("Total nil values", totalNil)
	fmt.Println("Total int values", totalInt)
	fmt.Println("Total string values", totalStr)
	fmt.Println("Total float64 values", totalFlt)
	fmt.Println("Total map values", totalMap)
	fmt.Println("Total slice values", totalSlc)
	fmt.Println("Total array values", totalArr)
	fmt.Println("Total pointer values", totalPtr)
	fmt.Println("Total boolean values", totalBln)
	fmt.Println("Total rune values", totalRue)
	fmt.Println("Total other values", totalOtr)
}

func main() {
	assertobject(uint(6), 5, "hi!", [3]int{3, 5, 8}, nil, 4.12, map[string]int{"j": 5, "k": 8})
	//assertobject([3]int{3, 5, 8}, []bool{false, true})
}
