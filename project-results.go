package reporter

type ProjectQueryResponse struct {
	Meta    Meta           `json:"meta"`
	Results []SearchResult `json:"results"`
}

type Meta struct {
	Search_id           string     `json:"search_id"`
	Total               int        `json:"total"`
	Offset              int        `json:"offset"`
	Limit               int        `json:"limit"`
	Sort_field          string     `json:"sort_field"`
	Sort_order          string     `json:"sort_order"`
	Sorted_by_relevance bool       `json:"sorted_by_relevance"`
	Properties          Properties `json:"properties"`
}

type Properties struct {
	URL string `json:"URL"`
}

type SearchResult struct {
	Appl_id            int               `json:"appl_id"`
	Subproject_id      string            `json:"subproject_id"`
	Fiscal_year        int               `json:"fiscal_year"`
	Project_num        string            `json:"project_num"`
	Project_serial_num string            `json:"project_serial_num"`
	Organization       ProjROrganization `json:"organization"`
	Award_type         string            `json:"award_type"`
	Activity_code      string            `json:"activity_code"`
	Award_amount       int               `json:"award_amount"`
	Is_active          bool              `json:"is_active"`
	Project_num_split  struct {
		Appl_type_code string `json:"appl_type_code"`
		Activity_code  string `json:"activity_code"`
		Ic_code        string `json:"ic_code"`
		Serial_num     string `json:"serial_num"`
		Support_year   string `json:"support_year"`
		Suffix_code    string `json:"suffix_code"`
	} `json:"project_num_split"`
	Principal_investigators []struct {
		Profile_id    int    `json:"profile_id"`
		First_name    string `json:"first_name"`
		Middle_name   string `json:"middle_name"`
		Last_name     string `json:"last_name"`
		Is_contact_pi bool   `json:"is_contact_pi"`
		Full_name     string `json:"full_name"`
		Title         string `json:"title"`
	} `json:"principal_investigators"`
	Contact_pi_name  string           `json:"contact_pi_name"`
	Program_officers []ProgramOfficer `json:"program_officers"`
	Agency_ic_admin  struct {
		Code         string `json:"code"`
		Abbreviation string `json:"abbreviation"`
		Name         string `json:"name"`
	} `json:"agency_ic_admin"`
	Agency_ic_fundings []struct {
		Fy           int     `json:"fy"`
		Code         string  `json:"code"`
		Name         string  `json:"name"`
		Abbreviation string  `json:"abbreviation"`
		Total_cost   float32 `json:"total_cost"`
	} `json:"agency_ic_fundings"`
	Cong_dist           string `json:"cong_dist"`
	Spending_categories struct {
		Values    []int `json:"values"`
		Match_all bool  `json:"match_all"`
	} `json:"spending_categories"`
	Project_start_date string `json:"project_start_date"`
	Project_end_date   string `json:"project_end_date"`
	Organization_type  struct {
		Name     string `json:"name"`
		Code     string `json:"code"`
		Is_other bool   `json:"is_other"`
	} `json:"organization_type"`
	Opportunity_number string `json:"opportunity_number"`
	Full_study_section struct {
		Srg_code            string `json:"srg_code"`
		Srg_flex            string `json:"srg_flex"`
		Sra_designator_code string `json:"sra_designator_code"`
		Sra_flex_code       string `json:"sra_flex_code"`
		Group_code          string `json:"group_code"`
		Name                string `json:"name"`
	} `json:"full_study_section"`
	Award_notice_date        string `json:"award_notice_date"`
	Is_new                   bool   `json:"is_new"`
	Mechanism_code_dc        string `json:"mechanism_code_dc"`
	Core_project_num         string `json:"core_project_num"`
	Terms                    string `json:"terms"`
	Pref_terms               string `json:"pref_terms"`
	Abstract_text            string `json:"abstract_text"`
	Project_title            string `json:"project_title"`
	Phr_text                 string `json:"phr_text"`
	Spending_categories_desc string `json:"spending_categories_desc"`
	Agency_code              string `json:"agency_code"`
	Covid_response           string `json:"covid_response"`
	Arra_funded              string `json:"arra_funded"`
	Budget_start             string `json:"budget_start"`
	Budget_end               string `json:"budget_end"`
	Cfda_code                string `json:"cfda_code"`
	Funding_mechanism        string `json:"funding_mechanism"`
	Direct_cost_amt          int    `json:"direct_cost_amt"`
	Indirect_cost_amt        int    `json:"indirect_cost_amt"`
	Project_detail_url       string `json:"project_detail_url"`
	Date_added               string `json:"date_added"`
}

type ProjROrganization struct {
	Org_name          string   `json:"org_name"`
	City              string   `json:"city"`
	Country           string   `json:"country"`
	Org_city          string   `json:"org_city"`
	Org_country       string   `json:"org_country"`
	Org_state         string   `json:"org_state"`
	Org_state_name    string   `json:"org_state_name"`
	Dept_type         string   `json:"dept_type"`
	Fips_country_code string   `json:"fips_country_code"`
	Org_duns          []string `json:"org_duns"`
	Org_ueis          []string `json:"org_ueis"`
	Primary_duns      string   `json:"primary_duns"`
	Primary_uei       string   `json:"primary_uei"`
	Org_fips          string   `json:"org_fips"`
	Org_ipf_code      string   `json:"org_ipf_code"`
	Org_zipcode       string   `json:"org_zipcode"`
	External_org_id   int      `json:"external_org_id"`
}

type ProgramOfficer struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	FullName   string `json:"full_name"`
}
