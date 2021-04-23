package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ip *IPAddr) String() string {
	var ipByteStrList []string
	for _, ipByte := range ip {
		ipByteStrList = append(ipByteStrList, fmt.Sprintf("%v", ipByte))
	}
	return strings.Join(ipByteStrList, ".")
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
