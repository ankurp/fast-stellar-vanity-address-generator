package main

import (
	"flag"
	"fmt"
	"runtime"

	"strings"

	"github.com/stellar/go/keypair"
)

func main() {
	prefix := flag.Bool("prefix", false, "a bool")
	flag.Parse()
	suffixes := flag.Args()
	for index, suffix := range suffixes {
		suffixes[index] = strings.ToUpper(suffix)
	}

	result := make(chan (*keypair.Full))
	for i := 0; i < runtime.NumCPU(); i++ {
		go search(suffixes, result, prefix)
	}

	keypair := <-result
	fmt.Println("Public: ", keypair.Address())
	fmt.Println("Secret: ", keypair.Seed())
}

func search(suffixes []string, result chan<- (*keypair.Full), prefix *bool) {
	var checkFunc func(string, string) bool
	if *prefix {
		checkFunc = strings.HasPrefix
	} else {
		checkFunc = strings.HasSuffix
	}
	for {
		keypair, err := keypair.Random()
		if err != nil {
			fmt.Println("Error generating Stellar Account", err)
			continue
		}

		for _, suffix := range suffixes {
			if checkFunc(keypair.Address(), suffix) {
				result <- keypair
			}
		}
	}
}
