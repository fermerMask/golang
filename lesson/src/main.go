package main

import (
	"fmt"
	"time"
)

// Hello world
/*
複数行のコメント
*/

func main() {
	fmt.Println("hello world")
	fmt.Println(time.Now())
	var i int = 100
	fmt.Println(i)
	var s string = "hello go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 100
		s2 string = "hello"
	)
	fmt.Println(i2, s2)
	var fl64 float64 = 2.4
	fmt.Println((fl64))

	var arr1 [3]int
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)

	fmt.Println(arr2[0])
	fmt.Println(len(arr1))

	var x interface{} //あらゆる型と互換性がある．
	fmt.Println(x)

	x = 1
	fmt.Println(x)

}
