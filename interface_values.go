//go:build ignore

package main

import "fmt"

type I interface {
	Mogu()
}

type T struct {
	Str string
}

// 型Tのポインタ変数をレシーバに持つメソッドMogu()があるので
// 型TはインターフェースIを実装していることになる
func (t *T) Mogu() {
	fmt.Println(t.Str)
}

type F float64

// 同じく型Fをレシーバに持つメソッドMogu()があるので
// 型FはインターフェースIを実装していることになる
func (f F) Mogu() {
	fmt.Println(f)
}

func main() {
	/*
		以下のように2行で書こうとすると、
		var i I
		i = &T{"hello"}

		以下のような警告が出たので、一行にした。
		should merge variable declaration with assignment on next line (S1021)
		https://zenn.dev/axpensive/articles/4cfbe6e1a9997a
	*/
	// 型TはインターフェースIを実装しているので、I型の変数にはTのポインタを代入できる
	var i I = &T{"hello"}
	describe(i)
	i.Mogu()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
describe(i)の出力は以下のように返ってくる。
(&{hello}, *main.T)

つまり、interfaceって実態がなさそうに思えるけど、
ちゃんとvalueとtypeがあるんだよって話。
この場合、valueは&{hello}、typeは*main.T
*/
