package main

import (
	"fmt"
	"log"

	"github.com/vamsi/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("vamsi")
	if err != nil {
		log.Fatalf("couldn't set current user: %v\n", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	fmt.Printf("Read config again: %+v\n", cfg)
}
