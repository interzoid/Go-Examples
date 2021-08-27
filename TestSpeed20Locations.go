// Get API key at www.interzoid.com/register 
package main;

import (
	"encoding/json"
	"fmt"
	"net/http"
	url2 "net/url"
)

// change to your API key here
const apiKey string = "YOURAPIKEYHERE"

// any URL to test the performance of from various global locations
const siteToTest string = "www.google.com"

// data structure to parse JSON returning from API into
type Payload struct {
	Origin string
	Seconds string
	PageResponseStatus string
	Code string
	Credits  int
	Contents string
}

func main() {

	//initialization
	thePayload := Payload{}
	loc := ""

  	// try each of the twenty locations
	for i := 1; i <= 20; i++ {
		switch i {
		case 1:
			loc = "Singapore"
		case 2:
			loc = "Frankfurt"
		case 3:
			loc = "Sydney"
		case 4:
			loc = "Tokyo"
		case 5:
			loc = "London"
		case 6:
			loc = "Stockholm"
		case 7:
			loc = "Mumbai"
		case 8:
			loc = "California"
		case 9:
			loc = "Virginia"
		case 10:
			loc = "SaoPaulo"
		case 11:
			loc = "Paris"
		case 12:
			loc = "Montreal"
		case 13:
			loc = "Milan"
		case 14:
			loc = "Seoul"
		case 15:
			loc = "Hong Kong"
		case 16:
			loc = "Ireland"
		case 17:
			loc = "Cape Town"
		case 18:
			loc = "Bahrain"
		case 19:
			loc = "Ohio"
		case 20:
			loc = "Osaka"
		}

		// call the API
		url := "https://api.interzoid.com/globalpageload?license="+apiKey+"&origin=" + url2.QueryEscape(loc) + "&url=" + siteToTest
		response, _ := http.Get(url)

		// check to ensure valid status code was received
		if response.StatusCode != 200 {
			fmt.Println("Error from:", loc, response.Status)

		// show result
		} else {
			json.NewDecoder(response.Body).Decode(&thePayload)
			fmt.Println(thePayload.Origin + " -> " + thePayload.Seconds)
		}

	}

}
