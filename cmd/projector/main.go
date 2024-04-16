package main

import (
	"fmt"
	"log"

	"github.com/nabinkatwal7/projector-go/pkg/projector"
)

func main() {
	opts, err := projector.GetOpts()

	if err != nil {
		log.Fatal("Unable to get options: %v", err)
	}

	fmt.Printf("opts: %+v\n", opts)
}