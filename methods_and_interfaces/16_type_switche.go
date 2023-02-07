//go:build ignore

package main

import "fmt"

func main() {
	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v \n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I do not know about type %T!\n", v)
	}
}

// 通常のswitchは値を指定するが、型switchは型で判定する。
