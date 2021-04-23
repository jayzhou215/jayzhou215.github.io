package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ip *IPAddr) String() string {
	var ipByteStrs []string
	for _, bytee := range ip {
		ipByteStrs = append(ipByteStrs, fmt.Sprintf("%v", bytee))
	}
	return strings.Join(ipByteStrs, ".")
}

type Data struct {
	Key string
}

func (data *Data) String() string {
	return data.Key
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, &ip)
	}

	data := &Data{"hello"}
	fmt.Println(data)
}
