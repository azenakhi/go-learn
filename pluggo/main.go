package main

import (
	"log"

	_ "github.com/azenakhi/go-learn/pluggo/ext"
	"github.com/azenakhi/go-learn/pluggo/plugin"
)

func main() {
	for _, v := range plugin.Registry {
		log.Println(v.Message())
	}
}
