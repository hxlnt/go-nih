package main

import (
	"fmt"

	reporter "github.com/hxlnt/go-nih"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {

	p := message.NewPrinter(language.English)

	// Define the search criteria
	criteria := reporter.ProjectSearchCriteria{
		Use_relevance:           false,
		Include_active_projects: true,
		// Exclude_subprojects:     true,
		Activity_codes: []string{"R43", "R44"},
		Fiscal_years:   []int{2025, 2026},
	}

	// Build query
	myQuery := reporter.NewProjectQuery().
		Criteria(criteria).
		IncludeFields("ProjectTitle", "AwardAmount", "ProjectStartDate", "ProjectEndDate", "Organization", "AgencyIcAdmin").
		SortDescendingBy("award_amount").
		MaxResultsToReturn(5)

	// Run the search
	res, err := myQuery.Search()

	// Handle errors
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Do something with the results
	fmt.Println("Search Results:\n")
	for i, result := range res.Results {
		p.Printf("%d) %s (%s-%s)\n   $%d awarded to %s by %s\n\n", i+1, result.Project_title, result.Project_start_date[:4], result.Project_end_date[:4], result.Award_amount, result.Organization.Org_name, result.Agency_ic_admin.Name)
	}

}
