package main

import "fmt"

type IPaddr [4]byte

func (ip IPaddr) String() string {
	return fmt.Sprintf("%d,%d,%d,%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	addrs := map[string]IPaddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
