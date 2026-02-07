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
	age         string
	dateCreated time.Time
}

func (userData User) output() {
	fmt.Printf("my name is %v %v (born on %v)", userData.firstName, userData.lastName, userData.age)
}

func NewUser(fName string, lName string, bDay string) *User {
	userData := User{
		firstName:   fName,
		lastName:    lName,
		age:         bDay,
		dateCreated: time.Now(),
	}

	return &userData
}
func main() {

	firstName := InputReader("Please Enter Your Firstname: ")
	lastName := InputReader("Please Enter Your Lastname: ")
	birth := InputReader("Your Birth Day: YYYY/MM/DD: ")

	userData := NewUser(firstName, lastName, birth)

	userData.output()
}
