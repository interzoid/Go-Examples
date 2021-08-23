// Get API key at www.interzoid.com/register           
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Payload struct {
	Email           string
	Response        string
	Info            string
	IsGovernment    string
	IsEducational   string
	IsOrganizational string
	IsVulgar      string
	IsGeneric     string
	Geolocation   string
	Code          string
	Credits       int
}

func main() {
    // Create the struct variable to store the decoded JSON response
    thePayload := Payload{}

    // Call the API
    // Can error check here too replacing underscore with 'err'
    response, _ := http.Get("https://api.interzoid.com/getemailinfo?license=
                YOURAPILICENSEKEY&email=winston.smith@maths.ox.ac.uk")

    if response.StatusCode != 200 {
        // Report any HTTP Errors
        fmt.Println("Error: ", response.Status)

    } else {
        // Decode the JSON and print a comma-delimited response,
        // one of many ways to format the response
        _ = json.NewDecoder(response.Body).Decode(&thePayload)

        fmt.Println(thePayload.Email + "," + thePayload.Response +
		"," + thePayload.Info +
		"," + thePayload.IsGovernment +
                "," + thePayload.IsEducational +
		"," + thePayload.IsOrganizational +
		"," + thePayload.IsVulgar +
		"," + thePayload.IsGeneric +
		"," + thePayload.Geolocation)
    }
}
