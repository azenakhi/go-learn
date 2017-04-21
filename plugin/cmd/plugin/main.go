package main

import (
	"log"

	_ "github.com/azenakhi/go-learn/plugin/extension"
	"github.com/azenakhi/go-learn/plugin/plug"
)

func main() {
	for _, v := range plug.Registry {
		log.Println(v.Message())
	}
}
