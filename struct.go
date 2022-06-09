package main

import "fmt"

func main() {
	user := User{
		Name:  "Krishna",
		EmpId: 101,
	}

	fmt.Println(user.Name)

}

type User struct {
	Name  string
	EmpId int
}
