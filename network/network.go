package network

import (
	"Corona_Test/test"
	"os"

	"encoding/json"
	"log"
	"net/http"
)

func GetUpdate() []test.Test {
	req, err := http.NewRequest("GET", "https://user-api.coronatest.nl/uitslagen/", nil)
	if err != nil {
		log.Printf("Could not create new request: %v", err)
	}

	bearer := os.Getenv("BEARER")
	req.Header.Set("Authorization", bearer)

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