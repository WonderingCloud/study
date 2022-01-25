package main

import (
	"fmt"
	"mine/backtrack"
	"strings"
	"time"
)

func main() {

	board := [][]byte{
		{'.', '.', '9', '.', '.', '.', '1', '2', '.'},
		{'.', '.', '.', '.', '7', '.', '.', '.', '9'},
		{'5', '.', '.', '6', '.', '3', '.', '.', '.'},
		{'7', '1', '.', '5', '.', '.', '.', '.', '.'},
		{'9', '.', '.', '.', '.', '.', '.', '.', '2'},
		{'.', '.', '.', '.', '.', '6', '.', '3', '1'},
		{'.', '.', '.', '8', '.', '1', '.', '.', '3'},
		{'8', '.', '.', '.', '5', '.', '.', '.', '.'},
		{'.', '9', '6', '.', '.', '.', '8', '.', '.'},
	}

	t1 := time.Now()
	backtrack.SolveSudoku(board)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	for i := range board {
		fmt.Println(string(board[i]))
	}
}

type Test struct {
	A int
	B int
}

func isPalindrome(s string) bool {
	bs := []byte(strings.ToLower(s))
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !((bs[i] >= 'a' && bs[i] <= 'z') || (bs[i] >= '0' && bs[i] <= '9')) {
			i++
		}

		for i < j && !((bs[j] >= 'a' && bs[j] <= 'z') || (bs[j] >= '0' && bs[j] <= '9')) {
			j--
		}

		if i >= j {
			break
		}
		if bs[i] != bs[j] {
			return false
		}
		i++
		j--
	}
	return true
}

type CQueue struct {
	s1 []int
	s2 []int
}

func Constructor() CQueue {
	return CQueue{s1: make([]int, 0), s2: make([]int, 0)}
}

func (this *CQueue) AppendTail(value int) {
	this.s1 = append(this.s1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.s2) == 0 {
		l := len(this.s1)
		for l != 0 {
			this.s2 = append(this.s2, this.s1[l-1])
			this.s1 = this.s1[:l-1]
			l--
		}
	}

	if len(this.s2) != 0 {
		val := this.s2[len(this.s2)-1]
		this.s2 = this.s2[:len(this.s2)-1]
		return val
	}
	return -1
}
