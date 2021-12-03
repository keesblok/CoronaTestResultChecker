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

	req.Header.Add("Authorization", os.Getenv("BEARER"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Getting results from coronatest.nl went wrong: %v", err)
	}
	defer resp.Body.Close()

	result := &[]test.Test{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Printf("Decoding json went wrong: %s\n", err)
		return *result
	}

	return *result
}