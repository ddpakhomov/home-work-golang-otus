package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// isValidURL проверяет, является ли строка допустимым URL.
func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	return err == nil
}

// sendGetRequest отправляет GET-запрос на указанный URL.
func sendGetRequest(url string) {
	if !isValidURL(url) {
		fmt.Printf("Invalid URL: %s\n", url)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating GET request: %v\n", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
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

// sendPostRequest отправляет POST-запрос на указанный URL с указанными данными.
func sendPostRequest(url string, data string) {
	if !isValidURL(url) {
		fmt.Printf("Invalid URL: %s\n", url)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBufferString(data))
	if err != nil {
		fmt.Printf("Error creating POST request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
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
