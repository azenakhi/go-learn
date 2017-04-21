package main

import (
	"log"

	"github.com/azenakhi/go-learn/reflex"
)

func main() {
	p := reflex.Person{}
	log.Println(p.Reflect())
}
