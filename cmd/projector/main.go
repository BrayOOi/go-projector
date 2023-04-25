package main

import (
	"log"

	"github.com/BrayOOi/go-projector.git/pkg/config"
)

func main() {
	opts, err := config.GetOpts()
	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}
}
