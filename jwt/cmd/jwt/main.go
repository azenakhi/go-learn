package main

import (
	"log"

	"github.com/azenakhi/go-learn/jwt"
)

func main() {
	token, _ := jwt.GetToken()
	log.Println(token)
}
