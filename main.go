package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	startingTime := time.Now().UTC()

	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "copesc@gmail.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		endingTime := time.Now().UTC()
		fmt.Println(endingTime.Sub(startingTime))
		return
	}

	stats(email)
	endingTime := time.Now().UTC()
	fmt.Println(endingTime.Sub(startingTime))
}
