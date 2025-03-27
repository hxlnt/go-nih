package main

import (
	"fmt"

	reporter "../reporter"
)

func main() {
	// Define the search criteria
	criteria := reporter.ProjectSearchCriteria{
		Use_relevance:           false,
		Include_active_projects: true,
		Exclude_subprojects:     true,
		Activity_codes:          []string{"R43", "R44"},
		Fiscal_years:            []int{2025, 2026},
	}

	// Perform the search
	res, err := reporter.NewProjectQuery().
		Criteria(criteria).
		IncludeFields("ProjectTitle", "AwardAmount", "FiscalYear", "Organization", "AgencyIcAdmin").
		// ExcludeFields("").
		SortDescendingBy("award_amount").
		Offset(0).
		MaxResultsPerPage(5).
		Search()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Search Results:")
	for _, result := range res.Results {
		fmt.Printf("Title: %s, Amount: %d, Year: %d, Org: %s\n", result.Project_title, result.Award_amount, result.Fiscal_year, result.Organization.Org_name)
	}

}
