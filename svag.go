package main

import (
	"fmt"
	"os"
	"runtime"

	"strings"

	"github.com/stellar/go/keypair"
)

func main() {
	suffixes := os.Args[1:]
	for index, suffix := range suffixes {
		suffixes[index] = strings.ToUpper(suffix)
	}
	numCPUs := runtime.NumCPU()
	for i := 1; i < numCPUs; i++ {
		go search(suffixes)
	}
	search(suffixes)
}

func search(suffixes []string) {
	for {
		keypair, err := keypair.Random()

		if err != nil {
			fmt.Println("Error generating Stellar Account")
			os.Exit(1)
		}

		for _, suffix := range suffixes {
			if strings.HasSuffix(keypair.Address(), suffix) {
				fmt.Println("Public: ", keypair.Address())
				fmt.Println("Secret: ", keypair.Seed())
				os.Exit(0)
			}
		}
	}
}
