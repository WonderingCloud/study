package main

import "fmt"

func main() {
	if live() == nil {
		fmt.Println("AAAAAA")
	} else {
		fmt.Println("BBBBBB")
	}
}

type People interface {
	Speak(string) string
	Show()
}

func live() People {
	var stu *Stduent
	return stu
}

type Stduent struct{}

func (stu *Stduent) Show() {

}

func (stu Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

