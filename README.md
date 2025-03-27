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
	criteria := nih.ProjectSearchCriteria{
		Use_relevance:           false,
		Include_active_projects: true,
		Fiscal_years:            []int{2025, 2026},
	}

	// Build query
	myQuery := nih.NewProjectQuery().
		Criteria(criteria).
		IncludeFields(
			nih.ProjectTitle,
			nih.AwardAmount,
			nih.ProjectStartDate,
			nih.ProjectEndDate).
		SortDescendingBy(nih.AwardAmount).
		MaxResultsToReturn(15)

	// Run the search
	res, err := myQuery.Search()

	if err != nil {
	    // Handle errors
	}

	for i, result := range res.Results {
		// Do something with the results
	}

}
```