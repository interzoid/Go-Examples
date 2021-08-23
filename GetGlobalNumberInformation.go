// Get API key at www.interzoid.com/register
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Payload struct {
    Country     string
    Region      string
    PrimaryCity string
    Language1   string
    Language2   string
    Language3   string
    Mobile      string
    Wealth      string
    Code        string
}

func main() {
    // Create the struct variable to store the decoded JSON response
    thePayload := Payload{}

    // Call the API
    response, _ := http.Get("https://api.interzoid.com/getglobalnumberinfo?license=
                YOURAPILICENSEKEY&intlnumber=+4906979550")

    if response.StatusCode != 200 {
        // Report any HTTP Errors
        fmt.Println("Error: ", response.Status)

    } else {
        // Decode the JSON and print a comma-delimited response,
        // one of many ways to format the response
        _ = json.NewDecoder(response.Body).Decode(&thePayload)

        fmt.Println(
        thePayload.Country + "," +
        thePayload.Region + "," +
        thePayload.PrimaryCity + "," +
        thePayload.Language1 + "," +
        thePayload.Language2 + "," +
        thePayload.Language3 + "," +
        thePayload.Mobile + "," +
        thePayload.Wealth +
        "\r\n")
    }
}
