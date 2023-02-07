//go:build ignore

package main

import "fmt"

func main() {
	var i interface{} = "hello"
	/*
		確認だけど、interface{}は全ての型と互換性を持っている。
		つまり何にでもなれる。

		あとついでに、interface{} は any と同意なので、
		var i any = "hello"
		でもOK!
	*/

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)

	src := "unkounko"
	printIf(src)

	src2 := []string{
		"test1",
		"test2",
		"test3",
	}
	printIf(src2)
}

// こんなふうに型によって処理を変化させることもできる。まあswitch使って書いても良い。
func printIf(src interface{}) {
	if value, ok := src.(int); ok {
		fmt.Printf("src is integer. value: %d\n", value)
		return
	}

	if value, ok := src.(string); ok {
		fmt.Printf("src is string. value: %s\n", value)
		return
	}

	if value, ok := src.([]string); ok {
		value = append(value, "hoge")
		fmt.Printf("src is slice. value: %s\n", value)
		return
	}

	fmt.Println("src is unknown type")
}

/*
身近な例だと、mockとかでinterface{}が引数で使われている。
mockの場合確かに引数はなんでもいいもんね。

func (mr *MockBasioAccessTokenRepositoryMockRecorder) InsertAccessToken(ctx, tx, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAccessToken", reflect.TypeOf((*MockBasioAccessTokenRepository)(nil).InsertAccessToken), ctx, tx, accessToken)
}
*/

/*
DBへクエリを投げる*sql.Stmtなどでも、where句で使う値は型なんでもいいので以下のような書き方ができる。

values := make([]any, 4)
values[0] = systemAuID
values[1] = stepCount
values[2] = totalingDate
values[3] = clientType

_, err = stmt.Exec(values...)
*/
