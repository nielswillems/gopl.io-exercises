// Exercise 1.9 Fetch prints the content found at each specified URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const httpProtocol = "http://"
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpProtocol) {
			url = httpProtocol + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Print("Response status code:", resp.Status)
	}
}
