//go:build ignore

package main

import "fmt"

type Password string

type Person struct {
	Name     string
	Password Password
}

// こんなふうにPassword型にStringerインターフェースを実装すると、出力内容がマスキング(テープ貼って見えなくするイメージ)できる。
// ex. ログでパスワードを表示したくない時
func (p Password) String() string {
	return "xxxxx"
}

func main() {
	a := Person{
		Name:     "Kaguya",
		Password: "passpasspass",
	}
	fmt.Println(a)
	// 出力
	// {Kaguya xxxxx}
}
