package reporter

import (
	"testing"
)

func Test_ValidQuery(t *testing.T) {
	criteria := ProjectSearchCriteria{
		ProjectNums: []string{"5R21AI175731-02"},
	}

	// Build query
	myQuery := NewProjectQuery().
		Criteria(criteria).
		IncludeFields(
			ProjectTitle,
			AwardAmount,
			ProjectStartDate,
			ProjectEndDate,
			Organization,
			AgencyIcAdmin).
		SortDescendingBy(ProjectStartDate).
		Offset(0).
		MaxResultsToReturn(5)

	// Run the search
	res, err := myQuery.Search()

	// Handle errors
	if err != nil {
		t.Fatalf("Expected no error but got '%s'", err)
		return
	}

	if res.Results[0].ProjectTitle != "Pyroptotic Macrophages Traps Against Shigella Infection" {
		t.Fatalf("Expected project title to be 'Pyroptotic Macrophages Traps Against Shigella Infection' but got '%s'", res.Results[0].ProjectTitle)
	}
}

func Test_InvalidQuery(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic: %v", r)
			return
		} else {
			t.Fatalf("Should panic but did not.")
		}
	}()

	criteria := ProjectSearchCriteria{
		IncludeActiveProjects: true,
	}

	// Build query
	myQuery := NewProjectQuery().
		Criteria(criteria).
		SortAscendingBy(ProjectTitle).
		// Too many results (max allowed is 500)--API will return a 400 error.
		MaxResultsToReturn(6000)

	// Run the search
	res, err := myQuery.Search()

	// Handle errors
	if err == nil {
		t.Fatal("Expected an error but got none.", res)
		return
	}

}

func Test_InvalidQuery2(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic: %v", r)
			if rStr, ok := r.(string); ok && rStr != "EOF" {
				t.Fatalf("Expected 'EOF' but got '%v'", r)
			} else {
				return
			}
		} else {
			t.Fatalf("Should panic but did not.")
		}
	}()

	criteria := ProjectSearchCriteria{}

	// Build query
	myQuery := NewProjectQuery().
		Criteria(criteria).
		ExcludeFields(ProgramOfficers).
		// Offset cannot be less than 0--API will return a generic 500 error.
		Offset(-1).
		MaxResultsToReturn(1)

	// Run the search
	res, err := myQuery.Search()

	// Handle errors
	if err == nil {
		t.Fatal("Expected an error but got none.", res)
		return
	}

}
