//go:build ignore

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
実行結果

(<nil>, <nil>)
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x10224744c]

goroutine 1 [running]:
main.main()
        /Users/funasaka/go/src/ghec/study-mediba2/tour_of_go_note/methods_and_interfaces/13_nil_interface_values.go:14 +0x5c
exit status 2
*/

/*
nilインターフェースを標準出力しても値も型もない。

Goのpanicを追ってみた記事はここ
https://qiita.com/nnao45/items/b8edaf82ece4f8114ddb

panicの意味の復習
シグナルや、code、addr、pcの意味も書かれているのでたまに目を通してもいい記事。
*/
