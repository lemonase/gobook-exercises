// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"strings"
	"io"
	// "io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// add protocol if missing
		if !(strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")) {
			url = "https://" + url
		}

		// fetch url
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// ioutil read body into buffer and print
		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s\n", b)

		// copy directly from body to stdout -- no buffer
		io.Copy(os.Stdout, resp.Body)

		// print response status and header
		fmt.Printf("RESPONSE STATUS: [%s]\n", resp.Status)
		for k, v := range resp.Header {
			fmt.Printf("%-15s: %s\n", k, v)
		}
		fmt.Println()

	}
}

//!-
