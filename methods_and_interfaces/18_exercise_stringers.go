//go:build ignore

package main

import "fmt"

type IPAddr [4]byte

// こんな感じで実装すれば、IPAddrを出力する際にドット区切りで表示できる。
func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

/*
出力

loopback: 127.0.0.1
googleDNS: 8.8.8.8
*/
