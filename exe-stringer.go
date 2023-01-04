package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (a IPAddr) String() string {
	var k string
	for i, val := range a {
		if i == 0 {
			k = fmt.Sprintf("%v", val)
		} else {
			k = k + "." + fmt.Sprintf("%v", val)
		}
	}
	return k
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
