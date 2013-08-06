package main

import (
	"fmt"
	"flag"
	"hash/fnv"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)

	// Open file if passed in.
	var file *os.File = os.Stdin
	var err error
	if flag.NArg() > 0 {
		file, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Slurp the input.
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Hash the input and print the hash to stdout.
	h := fnv.New64a()
	h.Reset()
	h.Write(b)
	hashcode := h.Sum64()
	fmt.Println(hashcode)
}
