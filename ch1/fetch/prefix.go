// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if ! ( strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") ) {
			url = "https://" + url
		}
		fmt.Printf("%s\n", url)
	}
}

//!-
