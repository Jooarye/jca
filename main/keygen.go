package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getRandomBytes(length int) []byte {
	file, err := os.Open("/dev/urandom")

	if err != nil {
		fmt.Println("Couldn't open /dev/urandom!")
		os.Exit(-1)
	}

	buffer := make([]byte, length)

	n, err := file.Read(buffer)

	if n != length {
		fmt.Println("Unknown error!")
		os.Exit(-2)
	}

	if err != nil {
		fmt.Println("Couldn't read from file!")
	}

	_ = file.Close()

	return buffer
}

func main() {
	size := flag.Int("size", sizes[3], "Size of the key")
	out := flag.String("out", "none", "File for the key")

	flag.Parse()

	if !contains(sizes, *size) {
		fmt.Println("Please select a valid size: " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sizes)), ", "), "[]"))
		os.Exit(-3)
	}

	if !isFlagPassed("out") {
		fmt.Printf("JCA %s\n\n", version)
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Printf("Generating key of size %d bytes!\n", *size)

	randomBytes := getRandomBytes(*size)

	err := ioutil.WriteFile(*out, randomBytes, 0644)

	if err != nil {
		fmt.Println("Couldn't write file!")
	}
}