package main

import (
	"flag"
	"fmt"
	"github.com/mckeowbc/superstrings"
	"os"
)

func main() {
	var num = flag.Uint("min-len", 6, "Minimum string length to find")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename>\n", os.Args[0])
		os.Exit(-1)
	}

	filename := flag.Args()[0]
	fin, err := os.Open(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(-1)
	}

	defer fin.Close()

	var offset uint64 = 0

	stringer := superstrings.NewStringer([]string{`Arabic`}, *num)

	for {
		data := make([]byte, 1024)

		count, err := fin.Read(data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(-1)
		}

		strings := stringer.GetStrings(data, offset)

		for _, found_string := range strings {
			fmt.Println(found_string)
		}

		if count < cap(data) {
			break
		}

		offset += uint64(count)
	}
}
