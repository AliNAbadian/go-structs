package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	firstName   string
	lastName    string
	age         int
	dateCreated time.Time
}

func main() {
	var userData User

	userData.firstName = inputReader("Please Enter Your Firstname: ")

	fmt.Println(userData.firstName)
}
