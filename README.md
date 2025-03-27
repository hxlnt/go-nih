# go-nih
Go client for NIH Reporter API


## Usage

```go
package main

import (

	reporter "github.com/hxlnt/go-nih"

)

func main() {

	// Define the search criteria
	criteria := reporter.ProjectSearchCriteria{
		Use_relevance:           false,
		Include_active_projects: true,
		Activity_codes: []string{"R43", "R44"},
		Fiscal_years:   []int{2025, 2026},
	}

	// Build query
	myQuery := reporter.NewProjectQuery().
		Criteria(criteria).
		IncludeFields("ProjectTitle", "AwardAmount", "ProjectStartDate", "ProjectEndDate", "Organization", "AgencyIcAdmin").
		SortDescendingBy("award_amount").
		MaxResultsToReturn(10)

	// Run the search
	res, err := myQuery.Search()

	// Handle errors
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, result := range res.Results {
		// Do something with the results
	}

}
```