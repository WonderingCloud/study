package util

import (
	"fmt"
	"math/rand"
	"strconv"
)

func GenRandomArray(length int) {
	tempMap := make(map[int]struct{})
	res := ""
	for len(tempMap) < length {
		num := rand.Intn(10000)
		_, ok := tempMap[num]
		if !ok {
			tempMap[num] = struct{}{}
			res = res + strconv.Itoa(num) + ","
		}
	}
	fmt.Println(res[:len(res)-1])
}
