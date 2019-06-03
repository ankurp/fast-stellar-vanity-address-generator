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
	result := make(chan (*keypair.Full))

	for i := 0; i < numCPUs; i++ {
		go search(suffixes, result)
	}

	keypair := <-result
	fmt.Println("Public: ", keypair.Address())
	fmt.Println("Secret: ", keypair.Seed())
}

func search(suffixes []string, result chan<- (*keypair.Full)) {
	for {
		keypair, err := keypair.Random()
		if err != nil {
			fmt.Println("Error generating Stellar Account", err)
			continue
		}

		for _, suffix := range suffixes {
			if strings.HasSuffix(keypair.Address(), suffix) {
				result <- keypair
			}
		}
	}
}
