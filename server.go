package main

import (
	"net/http"
)

type User struct {
	name  string
	age   string
	email string
}

func apiListening() {
	const port = ":8080"

	http.ListenAndServe(port, nil)
}
