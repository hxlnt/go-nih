# go-nih
![](https://img.shields.io/badge/Pre--release-0.0.1-red) ![Test coverage](https://img.shields.io/badge/Test%20coverage-95%25-green)

A Go library for easier interactions with the [NIH RePORTER API](https://api.reporter.nih.gov). This is a working pre-release version; use at your own risk.

## Features
Don't have [all 200+ NIH activity codes](https://grants.nih.gov/funding/activity-codes) memorized? Not sure what a parameter means or what data type it should return? Can't remember when field names should be passed in snake case or Pascal case? Then you might enjoy some of these features.

- In-editor type checking and validation for complex query parameters
  <img width="800" alt="Screenshot 2025-03-28 at 4 55 44 PM" src="https://github.com/user-attachments/assets/0682270a-52ad-4bb9-aaa0-d97917e2759a" />
- In-editor documentation of predefined values (*e.g.,* activity codes) accompanied by links to more information on the NIH website
  <img width="800" alt="Screenshot 2025-03-28 at 4 45 35 PM" src="https://github.com/user-attachments/assets/2b8f54fc-d646-499c-8c4f-8384aca33b8f" />
- Added debugging information when the API returns an error due to invalid search criteria
- Consistent casing/formatting of field names in queries (sort, include, and exclude) and responses

Currently, only project search is implemented. Publication search may be implemented at a future time.

## Usage
See the [example](example) folder for a complete working example.

```go
	// Define the search criteria
	criteria := nih.ProjectSearchCriteria{
		IncludeActiveProjects: true,
		FiscalYears:           []int{2025, 2026},
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

## To-do 
 - Completion of in-editor documentation
 - Implementation of publication search
 - Consider whether to offer more granular control over `http.Client` or leave as is
 - First major release
