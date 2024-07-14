package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func sendGetRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("GET request to %s\n", url)
	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", string(body))
}

func sendPostRequest(url string, data string) {
	resp, err := http.Post(url, "text/plain", bytes.NewBufferString(data))
	if err != nil {
		fmt.Printf("Error making POST request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("POST request to %s with data: %s\n", url, data)
	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", string(body))
}

func main() {
	method := flag.String("method", "GET", "HTTP method to use (GET or POST)")
	url := flag.String("url", "", "URL of the server")
	data := flag.String("data", "", "Data to send with POST request")
	flag.Parse()

	if *url == "" {
		fmt.Println("URL is required")
		os.Exit(1)
	}

	switch *method {
	case "GET":
		sendGetRequest(*url)
	case "POST":
		sendPostRequest(*url, *data)
	default:
		fmt.Println("Unsupported method. Use GET or POST.")
		os.Exit(1)
	}
}
