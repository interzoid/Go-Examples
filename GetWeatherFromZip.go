// Get API key at www.interzoid.com/register
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Payload struct {
	City string
	State string
	TempF string
	TempC string
	Weather string
	WindMPH string
	WindDir string
	RelativeHumidity string
	VisibilityMiles string
	AirQualityIndex string
	AirQualityCode string
	AirQuality string
	Sunrise string
	Sunset string
	DaylightHours string
	DaylightMinutes string
	Code string
	Credits int
}

func main() {
    // Create the struct variable to store the decoded JSON response
    thePayload := Payload{}

    // Call the API
    // Can error check here too replacing underscore with 'err'
    response, _ := http.Get("https://api.interzoid.com/getweatherzipcode?license=
                YOURAPILICENSEKEY&zip=94111")

    if response.StatusCode != 200 {
        // Report any HTTP Errors
        fmt.Println("Error: ", response.Status)

    } else {
        // Decode the JSON and print a comma-delimited response,
        // one of many ways to format the response
        _ = json.NewDecoder(response.Body).Decode(&thePayload)

	// display some of the fields
        fmt.Println(thePayload.City + "," + thePayload.State +
		"," + thePayload.TempF +
		"," + thePayload.TempC +
                "," + thePayload.Weather)
    }
}
