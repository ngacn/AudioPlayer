package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var Uuidgen func(string) string

// 使用闭包初始化Uuidgen
// uuid第2位是1开头表示track uuid
// 2开头表示playlist uuid
func uuidgen() func(string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func(sig string) string { //基于随机数的UUID生成器，Linux默认的
		var i int32
		switch sig {
		case "track":
			i = 0x1000
		case "playlist":
			i = 0x2000
		default:
			i = 0x0000
		}
		return fmt.Sprintf("%x%x-%x-%x%x",
			r.Int31(), r.Int31(),
			(r.Int31()&0x0fff)|i, //Generates a 32-bit Hex number of the form ?xxx (? indicates the UUID type)
			r.Int31(), r.Int31())
	}
}

func init() {
	Uuidgen = uuidgen()
}
