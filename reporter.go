package reporter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	projectURL     = "https://api.reporter.nih.gov/v2/projects/search"
	publicationURL = "https://api.reporter.nih.gov/v2/publications/search"
)

func CheckResponseDecoding(r *http.Response, model interface{}) {
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		fmt.Println("Error! Response Body:", r.Body)
		if r.StatusCode == http.StatusInternalServerError {
			fmt.Println("Check that your search criteria are valid.")
		}
		os.Exit(0)
	}
}
