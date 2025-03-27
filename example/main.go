package main

import (
	nih "github.com/hxlnt/go-nih"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {

	p := message.NewPrinter(language.English)

	// Define the search criteria
	criteria := nih.ProjectSearchCriteria{
		Use_relevance:           false,
		Include_active_projects: true,
		Activity_codes:          []string{"R43", "R44"},
		Fiscal_years:            []int{2025, 2026},
	}

	// Build query
	myQuery := nih.NewProjectQuery().
		Criteria(criteria).
		IncludeFields(
			nih.ProjectTitle,
			nih.AwardAmount,
			nih.ProjectStartDate,
			nih.ProjectEndDate,
			nih.Organization,
			nih.AgencyIcAdmin).
		SortDescendingBy(nih.AwardAmount).
		MaxResultsToReturn(5)

	// Run the search
	res, err := myQuery.Search()

	// Handle errors
	if err != nil {
		p.Println("Error:", err)
		return
	}

	// Do something with the results
	p.Printf("Search results:\n")
	for i, result := range res.Results {
		p.Printf("%d) %s (%s-%s)\n   $%d awarded to %s by %s\n\n", i+1, result.Project_title, result.Project_start_date[:4], result.Project_end_date[:4], result.Award_amount, result.Organization.Org_name, result.Agency_ic_admin.Name)
	}

}
