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
			// if r != "500 Internal Server Error" {
			// t.Fatalf("Expected '500 Internal Server Error' but got '%v'", r)
			// } else {
			return
			// }
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
		// Too many results--API will return an error.
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
