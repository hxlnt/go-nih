package reporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewProjectQuery() *ProjectQuery {
	return &ProjectQuery{}
}

func (pq *ProjectQuery) Criteria(criteria ProjectSearchCriteria) *ProjectQuery {
	pq.PQ_Criteria = criteria
	return pq
}

func (pq *ProjectQuery) IncludeFields(fields ...string) *ProjectQuery {
	pq.PQ_Include_fields = fields
	return pq
}

func (pq *ProjectQuery) ExcludeFields(fields ...string) *ProjectQuery {
	pq.PQ_Exclude_fields = fields
	return pq
}

func (pq *ProjectQuery) SortAscendingBy(field string) *ProjectQuery {
	pq.PQ_Sort_field = field
	pq.PQ_Sort_order = "asc"
	return pq
}

func (pq *ProjectQuery) SortDescendingBy(field string) *ProjectQuery {
	pq.PQ_Sort_field = field
	pq.PQ_Sort_order = "desc"
	return pq
}

// Default is 50 if not specified. Maximum allowed value is 500.
func (pq *ProjectQuery) MaxResultsPerPage(limit int) *ProjectQuery {
	pq.PQ_Limit = limit
	return pq
}

// Minimum and default values are 0. Maximum value is 14,999 or value specified in Limit() or 50 if Limit() is not specified.
func (pq *ProjectQuery) Offset(page int) *ProjectQuery {
	pq.PQ_Offset = page
	return pq
}

func (pq *ProjectQuery) Search() (*ProjectQueryResponse, error) {
	jsonReq, err := json.Marshal(pq)
	if err != nil {
		return nil, err
	}
	fmt.Println("JSON Entry:", string(jsonReq))
	request, _ := http.NewRequest(http.MethodPost, "https://api.reporter.nih.gov/v2/projects/search", bytes.NewBuffer(jsonReq))
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

func (pq *ProjectQuery) GetNextPage() (*ProjectQueryResponse, error) {
	// TODO: Implement this
	return pq.Search()
}

func (pq *ProjectQuery) GetPreviousPage() (*ProjectQueryResponse, error) {
	//TODO: Implement this
	return pq.Search()
}

type ProjectQuery struct {
	PQ_Criteria       ProjectSearchCriteria `json:"criteria"`
	PQ_Include_fields []string              `json:"include_fields,omitempty"`
	PQ_Exclude_fields []string              `json:"exclude_fields,omitempty"`
	PQ_Offset         int                   `json:"offset,omitempty"`
	PQ_Limit          int                   `json:"limit,omitempty"`
	PQ_Sort_field     string                `json:"sort_field,omitempty"`
	PQ_Sort_order     string                `json:"sort_order,omitempty"`
}

type ProjectSearchCriteria struct {
	Use_relevance             bool               `json:"use_relevance,omitempty"`
	Fiscal_years              []int              `json:"fiscal_years,omitempty"`
	Include_active_projects   bool               `json:"include_active_projects,omitempty"`
	Pi_names                  []Name             `json:"pi_names,omitempty"`
	Po_names                  []Name             `json:"po_names,omitempty"`
	Org_names                 []string           `json:"org_names,omitempty"`
	Pi_profile_ids            []int              `json:"pi_profile_ids,omitempty"`
	Org_cities                []string           `json:"org_cities,omitempty"`
	Org_states                []string           `json:"org_states,omitempty"`
	Project_nums              []string           `json:"project_nums,omitempty"`
	Appl_ids                  []string           `json:"appl_ids,omitempty"`
	Org_countries             []string           `json:"org_countries,omitempty"`
	Project_num_split         ProjectNumSplit    `json:"project_num_split,omitempty"`
	Agencies                  []string           `json:"agencies,omitempty"`
	Is_agency_admin           bool               `json:"is_agency_admin,omitempty"`
	Is_agency_funding         bool               `json:"is_agency_funding,omitempty"`
	Activity_codes            []string           `json:"activity_codes,omitempty"`
	Award_types               []string           `json:"award_types,omitempty"`
	Dept_types                []string           `json:"dept_types,omitempty"`
	Cong_dists                []string           `json:"cong_dists,omitempty"`
	Opportunity_numbers       []string           `json:"opportunity_numbers,omitempty"`
	Spending_categories       SpendingCategories `json:"spending_categories,omitempty"`
	Project_start_date        ProjectDate        `json:"project_start_date,omitempty"`
	Project_end_date          ProjectDate        `json:"project_end_date,omitempty"`
	Organization_type         []string           `json:"organization_type,omitempty"`
	Full_study_sections       []StudySection     `json:"full_study_sections,omitempty"`
	Award_notice_date         ProjectDate        `json:"award_notice_date,omitempty"`
	Award_amount_range        AwardAmountRange   `json:"award_amount_range,omitempty"`
	Date_added                ProjectDate        `json:"date_added,omitempty"`
	Exclude_subprojects       bool               `json:"exclude_subprojects,omitempty"`
	Multi_pi_only             bool               `json:"multi_pi_only,omitempty"`
	Newly_added_projects_only bool               `json:"newly_added_projects_only,omitempty"`
	Sub_project_only          bool               `json:"sub_project_only,omitempty"`
	Funding_mechanism         []string           `json:"funding_mechanism,omitempty"`
	Covid_response            []string           `json:"covid_response,omitempty"`
}

type SpendingCategories struct {
	Values    []int `json:"values,omitempty"`
	Match_all bool  `json:"match_all,omitempty"`
}

type ProjectDate struct {
	From_date string `json:"from_date,omitempty"`
	To_date   string `json:"to_date,omitempty"`
}

type AwardAmountRange struct {
	Min_amount int `json:"min_amount,omitempty"`
	Max_amount int `json:"max_amount,omitempty"`
}

type StudySection struct {
	IrgCode           string            `json:"irgCode"`
	SrgCode           string            `json:"srgCode"`
	SrgFlex           string            `json:"srgFlex"`
	SraDesignatorCode string            `json:"sraDesignatorCode"`
	SraFlexCode       string            `json:"sraFlexCode"`
	GroupCode         string            `json:"groupCode"`
	Name              string            `json:"name"`
	URL               string            `json:"url"`
	CmteID            int               `json:"cmteId"`
	ClusterIrgCode    string            `json:"clusterIrgCode"`
	Properties        map[string]string `json:"properties"`
}

type Name struct {
	AnyName    string `json:"anyName"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
}

type ProjectNumSplit struct {
	ApplTypeCode    string `json:"applTypeCode"`
	ActivityCode    string `json:"activityCode"`
	IcCode          string `json:"icCode"`
	SerialNum       string `json:"serialNum"`
	SupportYear     string `json:"supportYear"`
	FullSupportYear string `json:"fullSupportYear"`
	SuffixCode      string `json:"suffixCode"`
}
