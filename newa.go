package main

import (
	"container/list"
	"fmt"
)

type Student struct {
	fullName string
	age      int
	username string
	password int
}



func main() {
	fmt.Println("Welcome!")
	students := list.New()
	for true{
		var command int
		fmt.Println("1. Login\n2.Registern\n3.Quit")
		fmt.Scan(&command)
		if command == 1{
			fmt.Println("Enter username")
			var name string
			var password string
			fmt.Scan(&name)
			fmt.Println("Enter password")
			fmt.Scan(&password)
			bool flag := true
			for i := students.Front(); i != nil; i = i.Next() {
				//if
			}
		}else if command == 2{
			fmt.Println("Enter your full name")
			var fullname string
			fmt.Scan(&fullname)
		}else if command == 3{
			break
		}

	}

}
