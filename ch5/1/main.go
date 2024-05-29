package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	content := msg.getMessage()
	cost := len(content) * 3
	return content, cost
}

type message interface {
	getMessage() string
}

// // // // // // //

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}
