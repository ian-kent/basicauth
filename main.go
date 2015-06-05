package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected username (and optional password)")
		os.Exit(1)
	}

	username := os.Args[1]
	var password string

	if len(os.Args) >= 3 {
		password = os.Args[2]
	}

	auth := username + ":" + password
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(auth)))
}
