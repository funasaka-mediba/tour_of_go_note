//go:build ignore

package main

import "fmt"

type I interface {
	Mogu()
}

type Totoro struct {
	Str string
}

func (t *Totoro) Mogu() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.Str)
}

func main() {
	var i I

	var t *Totoro
	i = t
	describe(i) // (<nil>, *main.Totoro)
	i.Mogu()    // <nil>

	i = &Totoro{"hello"}
	describe(i) // (&{hello}, *main.Totoro)
	i.Mogu()    // hello
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
