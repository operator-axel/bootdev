package main

import "fmt"

func main() {
	fname := "ch40s"
	lname := "bold"
	age := 31
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing, but unlikely"

	userLog := fmt.Sprintf("%v %v %v %v %v %v", fname, lname, age, messageRate, isSubscribed, message)

	fmt.Println(userLog)
}
