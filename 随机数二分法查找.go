package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(100) + 1
	var i int
	max := 100
	min := 1
	for i = (max + min) / 2; ; i = (max + min) / 2 {
		if i == a {
			fmt.Printf(" 通过二分法查找结果为%d\n", i)
			break
		} else if i < a {
			min = i + 1
		} else {
			max = i - 1
		}
		if max < a || min > a {
			fmt.Printf("在随机数为%d的时候出现异常\n", a)
		}
	}
}
