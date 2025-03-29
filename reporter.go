package reporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	projectURL     = "https://api.reporter.nih.gov/v2/projects/search"
	publicationURL = "https://api.reporter.nih.gov/v2/publications/search"
)

// An object containing search criteria and other parameters which sent to the RePORTER API in a POST request in order to search for projects.
func NewProjectQuery() *ProjectQuery {
	return &ProjectQuery{}
}

// Adds search parameters (as a criteria object) to a project query.
func (pq *ProjectQuery) Criteria(criteria ProjectSearchCriteria) *ProjectQuery {
	pq.PQ_Criteria = criteria
	return pq
}

// Fields to explicitly include in the returned results.
func (pq *ProjectQuery) IncludeFields(fields ...ProjFieldName) *ProjectQuery {
	for _, field := range fields {
		pq.PQ_IncludeFields = append(pq.PQ_IncludeFields, ProjQFieldName[field])
	}
	return pq
}

// Fields for which data should not be returned in the results.
func (pq *ProjectQuery) ExcludeFields(fields ...ProjFieldName) *ProjectQuery {
	for _, field := range fields {
		pq.PQ_ExcludeFields = append(pq.PQ_ExcludeFields, ProjQFieldName[field])
	}
	return pq
}

// Field name to sort results, low to high.
func (pq *ProjectQuery) SortAscendingBy(field ProjFieldName) *ProjectQuery {
	pq.PQ_SortField = ProjRFieldName[field]
	pq.PQ_SortOrder = "asc"
	return pq
}

// Field name to sort results, high to low.
func (pq *ProjectQuery) SortDescendingBy(field ProjFieldName) *ProjectQuery {
	pq.PQ_SortField = ProjRFieldName[field]
	pq.PQ_SortOrder = "desc"
	return pq
}

// Default is 50 if not specified. Maximum allowed value is 500; values greater than 500 will cause the API to return an error.
func (pq *ProjectQuery) MaxResultsToReturn(limit int) *ProjectQuery {
	pq.PQ_Limit = limit
	return pq
}

// Minimum and default values are 0.
func (pq *ProjectQuery) Offset(page int) *ProjectQuery {
	pq.PQ_Offset = page
	return pq
}

// Completes a search on a given query, returning the results and error objects.
func (pq *ProjectQuery) Search() (*ProjectQueryResponse, error) {
	jsonReq, err := json.Marshal(pq)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest(http.MethodPost, projectURL, bytes.NewBuffer(jsonReq))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	result, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	var responseObject ProjectQueryResponse
	checkResponseDecoding(result, &responseObject)
	return &responseObject, nil
}

func checkResponseDecoding(r *http.Response, model interface{}) {
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		fmt.Println("Error! Response Body:", r.Status)
		if r.StatusCode == http.StatusInternalServerError {
			fmt.Println("Check that your search criteria are valid.")
		}
		panic(err)
	}
}
