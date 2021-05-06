package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func (t *Student) SetName(name string) {
	t.name = name
}

func (t *Student) SetAge(age int) {
	t.age = age
}

func main() {
	var a *Student
	a = &Student{"aaa", 20, 10}

	a.SetName("bbb")
	a.SetAge(30)
	fmt.Println(a)
}
