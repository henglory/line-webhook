package handler

import "fmt"

func Webhook(b []byte) {
	fmt.Printf("GET %s\n", string(b))
}
