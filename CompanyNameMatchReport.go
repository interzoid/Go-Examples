// Get API key at www.interzoid.com/register

package main;
// Obtains a similarity key from the Interzoid API for each record in the CSV file,
// and then sorts them to line up records that share the same similarity key
// building a possible duplicate match report in the process

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	url2 "net/url"
	"sort"
)

// data structure to hold, then sort, then process company data
type typeSimKeySlice struct {
	company   string
	business  string
	simkey    string
}

// data structure to store parsed JSON from API call
type Payload struct {
	Simkey string
	Code string
	Credits  int
}

func main() {

	// initialize instances of the data structure
	thePayload := Payload{}
	appendValue := typeSimKeySlice{}

	// set up the slice to store the similarity keys received from the API
	var keysSlice []typeSimKeySlice

	// open the input CSV file
	ifile, err := os.Open("companies.csv")

	// check for problem opening file
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// read the input file
	r := csv.NewReader(bufio.NewReader(ifile))

	// store the contents of the input file
	result, err := r.ReadAll()

	// ensure contents of the input data
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("\nProcessing...")

	// loop through the input data to prepare for match report
	for value := range result {

		// Prepare the query URL for the API call
		url := "https://api.interzoid.com/getcompanymatchadvanced?license=YOURAPIKEY&company=" + url2.QueryEscape(result[value][0]) + "&algorithm=wide"

		// call the API
		res, err := http.Get(url)

		// if ok, parse returned JSON, otherwise exit
		if res.StatusCode == 200 {
			json.NewDecoder(res.Body).Decode(&thePayload)
		} else {
			fmt.Println("HTTP error when calling API:", res.StatusCode)
			os.Exit(1)
		}

		// call the API for each value
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// prepare values for appending
		appendValue.company = result[value][0]
		appendValue.business = result[value][1]
		appendValue.simkey = thePayload.Simkey

		// append to the slice for processing later
		keysSlice = append(keysSlice, appendValue)

	}

	// sort the Slice by Simkey some matching similarity keys line up next to each other
	sort.Slice(keysSlice, func(i, j int) bool { return keysSlice[i].simkey < keysSlice[j].simkey })

	// begin match report
	fmt.Println("\nMatch Report\n================\n")

	// generate report by scanning to check for sequentially matching keys
	inmatch := false        // flag to know if in a collection of matches or not

	for value := range keysSlice {

		// do consecutive similarity keys match? If so, print to report
		if (value != 0) && (keysSlice[value].simkey == keysSlice[value-1].simkey) {
			if !inmatch {
				fmt.Println(keysSlice[value-1])
				inmatch = true
			}
			fmt.Println(keysSlice[value])
		} else {
			// no longer in a match - if a match collection just finished, print a blank line
			if inmatch {
				fmt.Println()
			}
			inmatch = false
		}
	}

	fmt.Println()
	fmt.Println("Finished.")
}
