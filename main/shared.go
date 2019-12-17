package main

import "flag"

var sizes = []int{ 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144 }
var version = "v0.9.2-dev"

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}