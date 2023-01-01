package main

import (
	"log"

	btcd "github.com/btcsuite/btcd"
)

func main() {
	// create new client instance
	client, err := btcd.New(&btcd.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "127.0.0.1:8332",
		User:         "rpcUsername",
		Pass:         "rpcPassword",
	}, nil)
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}

}
