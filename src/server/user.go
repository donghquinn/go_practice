package main

import (
	"fmt"
	"log"
)

type User struct {
	name        string
	age         string
	gender      string
	email       string
	dateOfBirth string
}

func setUser(name string, age string, gender string, email string, dateOfBirth string) {
	var Users []User

	Users = []User{
		{name, age, gender, email, dateOfBirth},
	}

	fmt.Println(Users)

	log.Default()

	// handleRequests()
}
