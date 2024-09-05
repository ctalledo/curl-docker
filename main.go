package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/Microsoft/go-winio"
)

func main() {
	// Define command-line arguments
	pipePath := flag.String("pipe", "", "Path to the named pipe")
	url := flag.String("url", "", "URL endpoint to query")

	// Parse the command-line arguments
	flag.Parse()

	// Create a dialer for the named pipe
	dial := func(network, addr string) (net.Conn, error) {
		return winio.DialPipe(*pipePath, nil)
	}

	// Create an HTTP client with a custom transport that uses the named pipe dialer
	client := &http.Client{
		Transport: &http.Transport{
			Dial: dial,
		},
		Timeout: time.Second * 5,
	}

	// Make the HTTP GET request
	resp, err := client.Get(*url)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
