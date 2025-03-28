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

func NewProjectQuery() *ProjectQuery {
	return &ProjectQuery{}
}

func (pq *ProjectQuery) Criteria(criteria ProjectSearchCriteria) *ProjectQuery {
	pq.PQ_Criteria = criteria
	return pq
}

func (pq *ProjectQuery) IncludeFields(fields ...ProjFieldName) *ProjectQuery {
	for _, field := range fields {
		pq.PQ_IncludeFields = append(pq.PQ_IncludeFields, ProjQFieldName[field])
	}
	return pq
}

func (pq *ProjectQuery) ExcludeFields(fields ...ProjFieldName) *ProjectQuery {
	for _, field := range fields {
		pq.PQ_ExcludeFields = append(pq.PQ_ExcludeFields, ProjQFieldName[field])
	}
	return pq
}

func (pq *ProjectQuery) SortAscendingBy(field ProjFieldName) *ProjectQuery {
	pq.PQ_SortField = ProjRFieldName[field]
	pq.PQ_SortOrder = "asc"
	return pq
}

func (pq *ProjectQuery) SortDescendingBy(field ProjFieldName) *ProjectQuery {
	pq.PQ_SortField = ProjRFieldName[field]
	pq.PQ_SortOrder = "desc"
	return pq
}

// Default is 50 if not specified. Maximum allowed value is 500.
func (pq *ProjectQuery) MaxResultsToReturn(limit int) *ProjectQuery {
	pq.PQ_Limit = limit
	return pq
}

// Minimum and default values are 0.
func (pq *ProjectQuery) Offset(page int) *ProjectQuery {
	pq.PQ_Offset = page
	return pq
}

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
	CheckResponseDecoding(result, &responseObject)
	return &responseObject, nil
}

func CheckResponseDecoding(r *http.Response, model interface{}) {
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		fmt.Println("Error! Response Body:", r.Status)
		if r.StatusCode == http.StatusInternalServerError {
			fmt.Println("Check that your search criteria are valid.")
		}
		panic(err)
	}
}
