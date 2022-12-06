package main

import (
	"net/http"
)

// 요청으로 들어올 유저 정보

func apiListening() {
	const port = ":8080"

	http.ListenAndServe(port, nil)
}

// 유저 정보 할당
