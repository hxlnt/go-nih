# go-nih
![](https://img.shields.io/badge/Pre--release-0.0.1-red) ![Test coverage](https://img.shields.io/badge/Test%20coverage-74%25-yellowgreen)

A Go client for NIH RePorter API. 

## Features

- Improved type checking and/or error reporting when search queries are malformed.
- In-editor documentation of parameters improves the developer experience.
- Field names have been made consistent within query and response objects, eliminating a common source of bugs when working with the original API (*e.g.,* the API expects "AwardAmount" in some parts of the query object and "award_amount" in others in reference to the same award amount query field.)

Full API documentation can be found at the [NIH RePORTER home page](https://api.reporter.nih.gov).

Currently, only project search is implemented. Publication search may be implemented at a future time.

## Usage
See the [example](example) folder for a complete working example.

```go

	// Define the search criteria
	criteria := nih.ProjectSearchCriteria{
		IncludeActiveProjects: true,
		FiscalYears:            []int{2025, 2026},
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
```