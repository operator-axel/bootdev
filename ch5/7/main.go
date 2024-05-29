package main

import (
	"fmt"
)

func (e email) cost() int {
	var costPerChar int
	if e.isSubscribed {
		costPerChar = 2
	} else {
		costPerChar = 5
	}
	return len(e.body) * costPerChar
}

func (e email) format() string {
	subStatus := "Subscribed"
	if !e.isSubscribed {
		subStatus = "Not Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, subStatus)
}

type email struct {
	isSubscribed bool
	body         string
}
