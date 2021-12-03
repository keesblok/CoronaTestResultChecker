package network

import (
	"Corona_Test/test"
	"fmt"
	"os"
	"strings"

	"encoding/json"
	"log"
	"net/http"
)

func GetUpdate() []test.Test {
	req, err := http.NewRequest("GET", "https://user-api.coronatest.nl/uitslagen/", http.NoBody)
	if err != nil {
		log.Printf("Could not create new request: %v", err)
	}

	bearer := os.Getenv("BEARER")
	req.Header.Set("Authorization", bearer)
	log.Printf(formatRequest(req))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Getting results from coronatest.nl went wrong: %v", err)
	}
	defer resp.Body.Close()

	// Print the HTTP Status Code and Status Name
	log.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	result := &[]test.Test{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Printf("Decoding json went wrong: %s\n", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("Response was: %q", resp.Body)
	}

	return *result
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v”, r.Host"))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}