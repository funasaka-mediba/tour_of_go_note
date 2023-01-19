//go:build ignore

package main

import "fmt"

func main() {
	// 0個のメソッドを指定されたインターフェース型って、
	// anyで書き換えOKだよね？
	var i interface{}
	describe(i)

	// 普通に値を入れても問題ない。
	i = 42
	describe(i)

	// 好きな型の値を入れても問題ない！
	i = "hello"
	describe(i)

	/*
		出力
		(<nil>, <nil>)
		(42, int)
		(hello, string)
	*/

	// -----------------------------------------------

	// anyで書き換えてみる。
	var a any
	describe(a)

	a = 111
	describe(a)

	a = "hogehoge"
	describe(a)

	/*
		出力
		(<nil>, <nil>)
		(111, int)
		(hogehoge, string)
	*/
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
interface{}とかanyは空のインターフェースなので、好きな型の値を入れれる。
なので、汎用性の高い関数の引数などで用いられることが多い。
例えば、知らない人はいない標準ライブラリfmt.Printfでも、第二引数はanyが使われてる。

func Printf(format string, a ...any) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

脱線するけど、...anyとか3点ドットは１つ以上だったら何個でもいいよという意味である。
なので、fmt.Printf("%v\n", name)みたいに１つでもいいし、
fmt.Printf("%v, %v\n", first, second)みたいに二つでもいいよ。

*/
