package main

import "fmt"

func loop() {
	data := []int{1, 2, 3}

	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}

}

type Person struct {
	name string
	age  int
}

func iterateStruct() {
	persons := [3]Person{{"mazhar", 25}, {"a", 11}, {"b", 22}}

	for i := range persons {
		p := persons[i]
		fmt.Println(p)
	}
}
