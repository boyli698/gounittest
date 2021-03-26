package main

import "fmt"

func main() {
	fmt.Println("asdf")
}

func Print01(str string) {
	fmt.Println(str)
}

func Print02(str string) {
	fmt.Println(str)
}

func Print03(str string) {
	if str != "" {
		fmt.Println(str)
	} else {
		fmt.Println("no str")
	}
}

func Print04(str string) {
	fmt.Println(str)
}
