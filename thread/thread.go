package thread

import (
	"fmt"
	"sync"
)

type void struct{}

// 交替打印数字和字母
func PrintNumberAndLetter() {
	letter, number := make(chan void), make(chan void)
	go func() {
		i := 1
		for {
			<-number
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			letter <- void{}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 0
		letters := "abcdefghijklmnopqrstuvwxyz"
		for {
			<-letter
			if i == len(letters) {
				return
			}
			fmt.Print(string(letters[i]))
			i++
			fmt.Print(string(letters[i]))
			i++
			number <- void{}
		}
	}()

	number <- void{}
	wg.Wait()
}
