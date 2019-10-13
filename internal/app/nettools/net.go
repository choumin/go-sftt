package nettools

import (
	"time"
)
func hash(num int) int {
	ret := (num % 997) * (num % 97)
	max := 65535
	min := 1023
	if ret > max {
		ret -= max
	}
	if ret < min {
		ret += min
	}
	return ret
}
func GeneratePort() int {
	ts := time.Now().Unix()
	base := 100
	a := int(ts) / base
	return hash(a)
}