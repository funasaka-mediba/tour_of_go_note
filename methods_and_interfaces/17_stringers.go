//go:build ignore

package main

import "fmt"

/*
type Stringer interface {
	String() string
}

fmtパッケージのStringerインターフェースはとてもシンプル。
このString()関数に何かの型をレシーバとして持たせ、メソッド化してあげれば、
その型はStringerインターフェースを実装した(纏うイメージ)事になる。

これを実装すると、fmtで出力した際に、カスタマイズした書式で出力してくれる。
出力は以下のようになる。

Kaguya (17 years) Miyuki (17 years)
*/

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{
		Name: "Kaguya",
		Age:  17,
	}
	b := Person{
		Name: "Miyuki",
		Age:  17,
	}
	fmt.Println(a, b)
}

/*
他にも
GoStringer
Formatter
などのインターフェースで書式が作れる。
参考: https://qiita.com/tenntenn/items/453a09c4c6d7f580d0ab#gostringer%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%95%E3%82%A7%E3%83%BC%E3%82%B9

まあ普通ならfmt.Printf()を使った方がわかりやすいが、
特定の型を出力する場合のみマスキングしたい場合などに用いられる。
*/
