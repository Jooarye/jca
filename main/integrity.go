package main

import (
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func calculateIntegritySum(file string) string {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("Couldn't open file!")
		os.Exit(-1)
	}

	md5 := md5.New()
	sha256 := sha256.New()

	md5.Write(data)
	sha256.Write(data)

	return fmt.Sprintf("sha256: %x\nmd5: %x", sha256.Sum(nil), md5.Sum(nil))
}

func main() {
	in := flag.String("in", "none", "The input file")

	flag.Parse()

	if !isFlagPassed("in") {
		fmt.Printf("JCA %s\n\n", version)
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Println(calculateIntegritySum(*in))
}
