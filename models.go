package reporter

type ProjectQueryResponse struct {
	Meta    Meta           `json:"meta"`
	Results []SearchResult `json:"results"`
}

type Meta struct {
	SearchId          string     `json:"search_id"`
	Total             int        `json:"total"`
	Offset            int        `json:"offset"`
	Limit             int        `json:"limit"`
	SortField         string     `json:"sort_field"`
	SortOrder         string     `json:"sort_order"`
	SortedByRelevance bool       `json:"sorted_by_relevance"`
	Properties        Properties `json:"properties"`
}

type Properties struct {
	Url string `json:"URL"`
}

type SearchResult struct {
	ApplId           int               `json:"appl_id"`
	SubprojectId     string            `json:"subproject_id"`
	FiscalYear       int               `json:"fiscal_year"`
	ProjectNum       string            `json:"project_num"`
	ProjectSerialNum string            `json:"project_serial_num"`
	Organization     ProjROrganization `json:"organization"`
	AwardType        string            `json:"award_type"`
	ActivityCode     string            `json:"activity_code"`
	AwardAmount      int               `json:"award_amount"`
	// A project is active if the current date is prior to the date given by BudgetEnd.
	IsActive        bool `json:"is_active"`
	ProjectNumSplit struct {
		ApplTypeCode string `json:"appl_type_code"`
		ActivityCode string `json:"activity_code"`
		IcCode       string `json:"ic_code"`
		SerialNum    string `json:"serial_num"`
		SupportYear  string `json:"support_year"`
		SuffixCode   string `json:"suffix_code"`
	} `json:"project_num_split"`
	PrincipalInvestigators []struct {
		ProfileId   int    `json:"profile_id"`
		FirstName   string `json:"first_name"`
		MiddleName  string `json:"middle_name"`
		LastName    string `json:"last_name"`
		IsContactPi bool   `json:"is_contact_pi"`
		FullName    string `json:"full_name"`
		Title       string `json:"title"`
	} `json:"principal_investigators"`
	ContactPiName   string           `json:"contact_pi_name"`
	ProgramOfficers []ProgramOfficer `json:"program_officers"`
	AgencyIcAdmin   struct {
		Code         string `json:"code"`
		Abbreviation string `json:"abbreviation"`
		Name         string `json:"name"`
	} `json:"agency_ic_admin"`
	AgencyIcFundings []struct {
		Fy           int     `json:"fy"`
		Code         string  `json:"code"`
		Name         string  `json:"name"`
		Abbreviation string  `json:"abbreviation"`
		TotalCost    float32 `json:"total_cost"`
	} `json:"agency_ic_fundings"`
	CongDist           string `json:"cong_dist"`
	SpendingCategories struct {
		Values   []int `json:"values"`
		MatchAll bool  `json:"match_all"`
	} `json:"spending_categories"`
	ProjectStartDate string `json:"project_start_date"`
	ProjectEndDate   string `json:"project_end_date"`
	OrganizationType struct {
		Name    string `json:"name"`
		Code    string `json:"code"`
		IsOther bool   `json:"is_other"`
	} `json:"organization_type"`
	OpportunityNumber string `json:"opportunity_number"`
	FullStudySection  struct {
		SrgCode           string `json:"srg_code"`
		SrgFlex           string `json:"srg_flex"`
		SraDesignatorCode string `json:"sra_designator_code"`
		SraFlexCode       string `json:"sra_flex_code"`
		GroupCode         string `json:"group_code"`
		Name              string `json:"name"`
	} `json:"full_study_section"`
	AwardNoticeDate        string `json:"award_notice_date"`
	IsNew                  bool   `json:"is_new"`
	MechanismCodeDc        string `json:"mechanism_code_dc"`
	CoreProjectNum         string `json:"core_project_num"`
	Terms                  string `json:"terms"`
	PrefRerms              string `json:"pref_terms"`
	AbstractText           string `json:"abstract_text"`
	ProjectTitle           string `json:"project_title"`
	PhrText                string `json:"phr_text"`
	SpendingCategoriesDesc string `json:"spending_categories_desc"`
	AgencyCode             string `json:"agency_code"`
	CovidResponse          string `json:"covid_response"`
	ArraFunded             string `json:"arra_funded"`
	BudgetStart            string `json:"budget_start"`
	BudgetEnd              string `json:"budget_end"`
	CfdaCode               string `json:"cfda_code"`
	FundingMechanism       string `json:"funding_mechanism"`
	DirectCostAmt          int    `json:"direct_cost_amt"`
	IndirectCostAmt        int    `json:"indirect_cost_amt"`
	ProjectDetailUrl       string `json:"project_detail_url"`
	DateAdded              string `json:"date_added"`
}

type ProjROrganization struct {
	OrgName         string   `json:"org_name"`
	City            string   `json:"city"`
	Country         string   `json:"country"`
	OrgCity         string   `json:"org_city"`
	OrgCountry      string   `json:"org_country"`
	OrgState        string   `json:"org_state"`
	OrgStateName    string   `json:"org_state_name"`
	DeptType        string   `json:"dept_type"`
	FipsCountryCode string   `json:"fips_country_code"`
	OrgDuns         []string `json:"org_duns"`
	OrgUeis         []string `json:"org_ueis"`
	PrimaryDuns     string   `json:"primary_duns"`
	PrimaryUei      string   `json:"primary_uei"`
	OrgFips         string   `json:"org_fips"`
	OrgIpfCode      string   `json:"org_ipf_code"`
	OrgZipcode      string   `json:"org_zipcode"`
	ExternalOrgId   int      `json:"external_org_id"`
}

type ProgramOfficer struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	FullName   string `json:"full_name"`
}

type ProjectQuery struct {
	PQ_Criteria      ProjectSearchCriteria `json:"criteria"`
	PQ_IncludeFields []string              `json:"include_fields,omitempty"`
	PQ_ExcludeFields []string              `json:"exclude_fields,omitempty"`
	PQ_Offset        int                   `json:"offset,omitempty"`
	PQ_Limit         int                   `json:"limit,omitempty"`
	PQ_SortField     string                `json:"sort_field,omitempty"`
	PQ_SortOrder     string                `json:"sort_order,omitempty"`
}

type ProjectSearchCriteria struct {
	UseRelevance           bool                   `json:"use_relevance,omitempty"`
	FiscalYears            []int                  `json:"fiscal_years,omitempty"`
	IncludeActiveProjects  bool                   `json:"include_active_projects,omitempty"`
	PiNames                []Name                 `json:"pi_names,omitempty"`
	PoNames                []Name                 `json:"po_names,omitempty"`
	OrgNames               []string               `json:"org_names,omitempty"`
	PiProfileIds           []int                  `json:"pi_profile_ids,omitempty"`
	OrgCities              []string               `json:"org_cities,omitempty"`
	OrgStates              []string               `json:"org_states,omitempty"`
	ProjectNums            []string               `json:"project_nums,omitempty"`
	ApplIds                []string               `json:"appl_ids,omitempty"`
	OrgCountries           []string               `json:"org_countries,omitempty"`
	ProjectNumSplit        ProjectNumSplitType    `json:"project_num_split,omitempty"`
	Agencies               []string               `json:"agencies,omitempty"`
	IsAgencyAdmin          bool                   `json:"is_agency_admin,omitempty"`
	IsAgencyFunding        bool                   `json:"is_agency_funding,omitempty"`
	ActivityCodes          []Activity             `json:"activity_codes,omitempty"`
	AwardTypes             []string               `json:"award_types,omitempty"`
	DeptTypes              []string               `json:"dept_types,omitempty"`
	CongDists              []string               `json:"cong_dists,omitempty"`
	OpportunityNumbers     []string               `json:"opportunity_numbers,omitempty"`
	SpendingCategories     SpendingCategoriesType `json:"spending_categories,omitempty"`
	ProjectStartDate       ProjectDate            `json:"project_start_date,omitempty"`
	ProjectEndDate         ProjectDate            `json:"project_end_date,omitempty"`
	OrganizationType       []string               `json:"organization_type,omitempty"`
	FullStudySections      []StudySection         `json:"full_study_sections,omitempty"`
	AwardNoticeDate        ProjectDate            `json:"award_notice_date,omitempty"`
	AwardAmountRange       AwardAmountRange       `json:"award_amount_range,omitempty"`
	DateAdded              ProjectDate            `json:"date_added,omitempty"`
	ExcludeSubprojects     bool                   `json:"exclude_subprojects,omitempty"`
	MultiPiOnly            bool                   `json:"multi_pi_only,omitempty"`
	NewlyAddedProjectsOnly bool                   `json:"newly_added_projects_only,omitempty"`
	SubProjectOnly         bool                   `json:"sub_project_only,omitempty"`
	FundingMechanism       []string               `json:"funding_mechanism,omitempty"`
	CovidResponse          []string               `json:"covid_response,omitempty"`
}

type SpendingCategoriesType struct {
	Values   []int `json:"values,omitempty"`
	MatchAll bool  `json:"match_all,omitempty"`
}

type ProjectDate struct {
	FromDate string `json:"from_date,omitempty"`
	ToDate   string `json:"to_date,omitempty"`
}

type AwardAmountRange struct {
	MinAmount int `json:"min_amount,omitempty"`
	MaxAmount int `json:"max_amount,omitempty"`
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

type ProjectNumSplitType struct {
	ApplTypeCode    string `json:"applTypeCode"`
	ActivityCode    string `json:"activityCode"`
	IcCode          string `json:"icCode"`
	SerialNum       string `json:"serialNum"`
	SupportYear     string `json:"supportYear"`
	FullSupportYear string `json:"fullSupportYear"`
	SuffixCode      string `json:"suffixCode"`
}

type Activity int

const (
	// Research Facilities Construction Grant: https://grants.nih.gov/funding/activity-codes/C06
	C06 Activity = iota
	// International Research Training Grant: https://grants.nih.gov/funding/activity-codes/
	D43
	D71
	DP1
	DP2
	DP3
	DP4
	DP5
	DP7
	E11
	F05
	F30
	F31
	F32
	F33
	F37
	F38
	F99
	FI2
	FM1
	G07
	G08
	G11
	G12
	G13
	G20
	G94
	H13
	H23
	H25
	H28
	H50
	H57
	H62
	H64
	H75
	H79
	HD4
	I01
	I80
	IK3
	K00
	K01
	K02
	K05
	K06
	K07
	K08
	K12
	K14
	K18
	K21
	K22
	K23
	K24
	K25
	K26
	K30
	K32
	K38
	K43
	K76
	K99
	KD1
	KL1
	KL2
	KM1
	L30
	L32
	L40
	L50
	L60
	L70
	M01
	OT2
	P01
	P20
	P2C
	P30
	// Animal (Mammalian and Nonmammalian) Model, and Animal and Biological Material Resource Grant: https://grants.nih.gov/funding/activity-codes/P40
	P40
	// Biotechnology Resource Grant: https://grants.nih.gov/funding/activity-codes/P41
	P41
	// Harzardous Substance Basic Research Grant: https://grants.nih.gov/funding/activity-codes/
	P42
	// Specialized Center: https://grants.nih.gov/funding/activity-codes/P50
	P50
	// Primate Research Center Grant: https://grants.nih.gov/funding/activity-codes/P51
	P51
	// Comprehensive Center: https://grants.nih.gov/funding/activity-codes/P60
	P60
	// Linked Center Core Grant: https://grants.nih.gov/funding/activity-codes/PL1
	PL1
	// Program Project or Center with Complex Structure: https://grants.nih.gov/funding/activity-codes/PM1
	PM1
	// Concept Development Award: https://grants.nih.gov/funding/activity-codes/PN1
	PN1
	// Research Development Center: https://grants.nih.gov/funding/activity-codes/PN2
	PN2
	// Research Transition Award: https://grants.nih.gov/funding/activity-codes/R00
	R00
	// Research Project Grant: https://grants.nih.gov/funding/activity-codes/R01
	R01
	// Small Research Grant: https://grants.nih.gov/funding/activity-codes/R03
	R03
	// Conference: https://grants.nih.gov/funding/activity-codes/R13
	R13
	// Research Enhancement Award: https://grants.nih.gov/funding/activity-codes/R15
	R15
	// Research Excellence Award: https://grants.nih.gov/funding/activity-codes/R16
	R16
	// Research Demonstration and Dissemination Projects: https://grants.nih.gov/funding/activity-codes/R18
	R18
	// Exploratory/Developmental Grant: https://grants.nih.gov/funding/activity-codes/R21
	R21
	// Resource-Related Research Project (Biomedical): https://grants.nih.gov/funding/activity-codes/R24
	R24
	// Education Project: https://grants.nih.gov/funding/activity-codes/R25
	R25
	// Resource-Related Research Project (Clinical): https://grants.nih.gov/funding/activity-codes/R28
	R28
	R2F
	R30
	R33
	R34
	R35
	R36
	R37
	R38
	R41
	R42
	R43
	R44
	R49
	R50
	R55
	R56
	R61
	R90
	RC1
	RC2
	RC3
	RC4
	RF1
	RL1
	RL2
	RL5
	RL9
	RM1
	RS1
	S06
	// Biomedical Research Support Grant: https://grants.nih.gov/funding/activity-codes/S07
	S07
	// Biomedical Research Support Shared Instrumentation Grant: https://grants.nih.gov/funding/activity-codes/S10
	S10
	S11
	S15
	S21
	S22
	SB1
	SC1
	SC2
	SC3
	SI2
	T01
	T02
	T09
	// Conferences: https://grants.nih.gov/funding/activity-codes/T14
	T14
	// Continuing Education Training Grant: https://grants.nih.gov/funding/activity-codes/T15
	T15
	// Institutional National Research Service Award: https://grants.nih.gov/funding/activity-codes/T32
	T32
	T34
	// NRSA Short-Term Research Training: https://grants.nih.gov/funding/activity-codes/T35
	T35
	// Minority International Research Training Grant: https://grants.nih.gov/funding/activity-codes/T37
	T37
	T42
	T90
	// Linked Training Award: https://grants.nih.gov/funding/activity-codes/TL1
	TL1
	TL4
	TU2
	U01
	U09
	U10
	U11
	U13
	U17
	U18
	U19
	U1A
	U1B
	U1Q
	U1V
	U21
	// HIV/STD Preventative Services: https://grants.nih.gov/funding/activity-codes/U22
	U22
	U23
	U24
	U27
	U2C
	U2F
	U2G
	U2R
	U30
	U32
	U34
	U36
	U3R
	U41
	U42
	// Small Business Innovation Research (SBIR) Cooperative Agreement Phase I: https://grants.nih.gov/funding/activity-codes/U43
	U43
	// Small Business Innovation Research (SBIR) Cooperative Agreement Phase II: https://grants.nih.gov/funding/activity-codes/U44
	U44
	U45
	U47
	U48
	U49
	U50
	U51
	U52
	U53
	U54
	U55
	// Exploratory Grants Cooperative Agreement: https://grants.nih.gov/funding/activity-codes/U56
	U56
	U57
	U58
	U59
	U60
	U61
	U62
	U65
	// Immunization Demonstration Projects Cooperative Agreement: https://grants.nih.gov/funding/activity-codes/U66
	U66
	U75
	U79
	U81
	U82
	U83
	U84
	U86
	U87
	U90
	UA1
	UA5
	// Commercialization Readiness Program Cooperative Agreement: https://grants.nih.gov/funding/activity-codes/UB1
	UB1
	UC1
	UC2
	UC3
	UC4
	UC6
	UC7
	UD1
	UE1
	UE2
	UE5
	UF1
	UG1
	UG3
	UH1
	UH2
	UH3
	// Hazmat Training at DOE Nuclear Weapons Complex: https://grants.nih.gov/funding/activity-codes/UH4
	UH4
	// Linked Specialized Center Cooperative Agreement: https://grants.nih.gov/funding/activity-codes/UL1
	UL1
	UM1
	UM2
	// Early Independence Award Cooperative Agreement: https://grants.nih.gov/funding/activity-codes/UP5
	UP5
	UR6
	UR8
	US3
	US4
	// Small Business Technology Transfer (STTR) Cooperative Agreement Phase I: https://grants.nih.gov/funding/activity-codes/UT1
	UT1
	// Small Business Technology Transfer (STTR) Cooperative Agreement Phase II: https://grants.nih.gov/funding/activity-codes/UT2
	UT2
	// Rape Prevention and Education Grant: https://grants.nih.gov/funding/activity-codes/VF1
	VF1
	X01
	X02
	X98
	X99
)

var ActivityCodeType = map[Activity]string{
	C06: "C06",
	D43: "D43",
	D71: "D71",
	DP1: "DP1",
	DP2: "DP2",
	DP3: "DP3",
	DP4: "DP4",
	DP5: "DP5",
	E11: "E11",
	F05: "F05",
	F30: "F30",
	F31: "F31",
	F32: "F32",
	F33: "F33",
	F37: "F37",
	F38: "F38",
	F99: "F99",
	FI2: "FI2",
	FM1: "FM1",
	G07: "G07",
	G08: "G08",
	G11: "G11",
	G12: "G12",
	G13: "G13",
	G20: "G20",
	G94: "G94",
	H13: "H13",
	H23: "H23",
	H25: "H25",
	H28: "H28",
	H50: "H50",
	H57: "H57",
	H62: "H62",
	H64: "H64",
	H75: "H75",
	H79: "H79",
	HD4: "HD4",
	I01: "I01",
	I80: "I80",
	IK3: "IK3",
	K00: "K00",
	K01: "K01",
	K02: "K02",
	K05: "K05",
	K06: "K06",
	K07: "K07",
	K08: "K08",
	K12: "K12",
	K14: "K14",
	K18: "K18",
	K21: "K21",
	K22: "K22",
	K23: "K23",
	K24: "K24",
	K25: "K25",
	K26: "K26",
	K30: "K30",
	K32: "K32",
	K38: "K38",
	K43: "K43",
	K76: "K76",
	K99: "K99",
	KD1: "KD1",
	KL1: "KL1",
	KL2: "KL2",
	KM1: "KM1",
	L30: "L30",
	L32: "L32",
	L40: "L40",
	L50: "L50",
	L60: "L60",
	L70: "L70",
	M01: "M01",
	OT2: "OT2",
	P01: "P01",
	P20: "P20",
	P2C: "P2C",
	P30: "P30",
	P40: "P40",
	P41: "P41",
	P42: "P42",
	P50: "P50",
	P51: "P51",
	P60: "P60",
	PL1: "PL1",
	PM1: "PM1",
	PN1: "PN1",
	PN2: "PN2",
	R00: "R00",
	R01: "R01",
	R03: "R03",
	R13: "R13",
	R15: "R15",
	R16: "R16",
	R18: "R18",
	R21: "R21",
	R24: "R24",
	// Education Project: https://grants.nih.gov/funding/activity-codes/R25
	R25: "R25",
	R28: "R28",
	R2F: "R2F",
	R30: "R30",
	R33: "R33",
	R34: "R34",
	R35: "R35",
	R36: "R36",
	R37: "R37",
	R38: "R38",
	R41: "R41",
	R42: "R42",
	R43: "R43",
	R44: "R44",
	R49: "R49",
	R50: "R50",
	R55: "R55",
	R56: "R56",
	R61: "R61",
	R90: "R90",
	RC1: "RC1",
	RC2: "RC2",
	RC3: "RC3",
	RC4: "RC4",
	RF1: "RF1",
	RL1: "RL1",
	RL2: "RL2",
	RL5: "RL5",
	RL9: "RL9",
	RM1: "RM1",
	RS1: "RS1",
	S06: "S06",
	// Biomedical Research Support Grant: https://grants.nih.gov/funding/activity-codes/S07
	S07: "S07",
	S10: "S10",
	S11: "S11",
	S15: "S15",
	S21: "S21",
	S22: "S22",
	SB1: "SB1",
	SC1: "SC1",
	SC2: "SC2",
	SC3: "SC3",
	SI2: "SI2",
	T01: "T01",
	T02: "T02",
	T09: "T09",
	T14: "T14",
	T15: "T15",
	T32: "T32",
	T34: "T34",
	T35: "T35",
	T37: "T37",
	T42: "T42",
	T90: "T90",
	TL1: "TL1",
	TL4: "TL4",
	TU2: "TU2",
	U01: "U01",
	U09: "U09",
	U10: "U10",
	U11: "U11",
	U13: "U13",
	U17: "U17",
	U18: "U18",
	U19: "U19",
	U1A: "U1A",
	U1B: "U1B",
	U1Q: "U1Q",
	U1V: "U1V",
	U21: "U21",
	U22: "U22",
	U23: "U23",
	U24: "U24",
	U27: "U27",
	U2C: "U2C",
	U2F: "U2F",
	U2G: "U2G",
	U2R: "U2R",
	U30: "U30",
	U32: "U32",
	U34: "U34",
	U36: "U36",
	U3R: "U3R",
	U41: "U41",
	U42: "U42",
	U43: "U43",
	U44: "U44",
	U45: "U45",
	U47: "U47",
	U48: "U48",
	U49: "U49",
	U50: "U50",
	U51: "U51",
	U52: "U52",
	U53: "U53",
	U54: "U54",
	U55: "U55",
	U56: "U56",
	U57: "U57",
	U58: "U58",
	U59: "U59",
	U60: "U60",
	U61: "U61",
	U62: "U62",
	U65: "U65",
	U66: "U66",
	U75: "U75",
	U79: "U79",
	U81: "U81",
	U82: "U82",
	U83: "U83",
	U84: "U84",
	U86: "U86",
	U87: "U87",
	U90: "U90",
	UA1: "UA1",
	UA5: "UA5",
	UB1: "UB1",
	UC1: "UC1",
	UC2: "UC2",
	UC3: "UC3",
	UC4: "UC4",
	UC6: "UC6",
	UC7: "UC7",
	UD1: "UD1",
	UE1: "UE1",
	UE2: "UE2",
	UE5: "UE5",
	UF1: "UF1",
	UG1: "UG1",
	UG3: "UG3",
	UH1: "UH1",
	UH2: "UH2",
	UH3: "UH3",
	UH4: "UH4",
	UL1: "UL1",
	UM1: "UM1",
	UM2: "UM2",
	UP5: "UP5",
	UR6: "UR6",
	UR8: "UR8",
	US3: "US3",
	US4: "US4",
	// Small Business Technology Transfer (STTR) Cooperative Agreement Phase I: https://grants.nih.gov/funding/activity-codes/UT1
	UT1: "UT1",
	// Small Business Technology Transfer (STTR) Cooperative Agreement Phase II: https://grants.nih.gov/funding/activity-codes/UT2
	UT2: "UT2",
	// Rape Prevention and Education Grant: https://grants.nih.gov/funding/activity-codes/VF1
	VF1: "VF1",
	X01: "X01",
	X02: "X02",
	X98: "X98",
	X99: "X99",
}

type ProjFieldName int

const (
	AbstractText ProjFieldName = iota
	ActivityCode
	AgencyCode
	AgencyIcAdmin
	AgencyIcFundings
	ApplId
	ArraFunded
	AwardAmount
	AwardNoticeDate
	AwardType
	BudgetEnd
	BudgetStart
	CfdaCode
	CongDist
	ContactPiName
	CoreProjectNum
	CovidResponse
	DateAdded
	DirectCostAmt
	IndirectCostAmt
	// A project is active if the current date is prior to the date given by BudgetEnd.
	IsActive
	IsNew
	FiscalYear
	FullStudySection
	FundingMechanism
	MechanismCodeDc
	OpportunityNumber
	Organization
	OrganizationType
	PhrText
	Pi
	PrefTerms
	ProgramOfficers
	ProjectDetailUrl
	ProjectEndDate
	ProjectNumSplit
	ProjectStartDate
	ProjectTitle
	SpendingCategories
	SpendingCategoriesDesc
	SubprojectId
	Terms
)

var ProjQFieldName = map[ProjFieldName]string{
	AbstractText:           "AbstractText",
	ActivityCode:           "ActivityCode",
	AgencyCode:             "AgencyCode",
	AgencyIcAdmin:          "AgencyIcAdmin",
	AgencyIcFundings:       "AgencyIcFundings",
	ApplId:                 "ApplId",
	ArraFunded:             "ArraFunded",
	AwardAmount:            "AwardAmount",
	AwardNoticeDate:        "AwardNoticeDate",
	AwardType:              "AwardType",
	BudgetEnd:              "BudgetEnd",
	BudgetStart:            "BudgetStart",
	CfdaCode:               "CfdaCode",
	CongDist:               "CongDist",
	ContactPiName:          "ContactPiName",
	CoreProjectNum:         "CoreProjectNum",
	CovidResponse:          "CovidResponse",
	DateAdded:              "DateAdded",
	DirectCostAmt:          "DirectCostAmt",
	IndirectCostAmt:        "IndirectCostAmt",
	IsActive:               "IsActive",
	IsNew:                  "IsNew",
	FiscalYear:             "FiscalYear",
	FullStudySection:       "FullStudySection",
	FundingMechanism:       "FundingMechanism",
	MechanismCodeDc:        "MechanismCodeDc",
	OpportunityNumber:      "OpportunityNumber",
	Organization:           "Organization",
	OrganizationType:       "OrganizationType",
	PhrText:                "PhrText",
	Pi:                     "Pi",
	PrefTerms:              "PrefTerms",
	ProgramOfficers:        "ProgramOfficers",
	ProjectDetailUrl:       "ProjectDetailUrl",
	ProjectEndDate:         "ProjectEndDate",
	ProjectNumSplit:        "ProjectNumSplit",
	ProjectStartDate:       "ProjectStartDate",
	ProjectTitle:           "ProjectTitle",
	SpendingCategories:     "SpendingCategories",
	SpendingCategoriesDesc: "SpendingCategoriesDesc",
	SubprojectId:           "SubprojectId",
	Terms:                  "Terms",
}

var ProjRFieldName = map[ProjFieldName]string{
	AbstractText:           "abstract_text",
	ActivityCode:           "activity_code",
	AgencyCode:             "agency_code",
	AgencyIcAdmin:          "agency_ic_admin",
	AgencyIcFundings:       "agency_ic_fundings",
	ApplId:                 "appl_id",
	ArraFunded:             "arra_funded",
	AwardAmount:            "award_amount",
	AwardNoticeDate:        "award_notice_date",
	AwardType:              "award_type",
	BudgetEnd:              "budget_end",
	BudgetStart:            "budget_start",
	CfdaCode:               "cfda_code",
	CongDist:               "cong_dist",
	ContactPiName:          "contact_pi_name",
	CoreProjectNum:         "core_project_num",
	CovidResponse:          "covid_response",
	DateAdded:              "date_added",
	DirectCostAmt:          "direct_cost_amt",
	IndirectCostAmt:        "indirect_cost_amt",
	IsActive:               "is_active",
	IsNew:                  "is_new",
	FiscalYear:             "fiscal_year",
	FullStudySection:       "full_study_section",
	FundingMechanism:       "funding_mechanism",
	MechanismCodeDc:        "mechanism_code_dc",
	OpportunityNumber:      "opportunity_number",
	Organization:           "organization",
	OrganizationType:       "organization_type",
	PhrText:                "phr_text",
	Pi:                     "pi",
	PrefTerms:              "pref_terms",
	ProgramOfficers:        "program_officers",
	ProjectDetailUrl:       "project_detail_url",
	ProjectEndDate:         "project_end_date",
	ProjectNumSplit:        "project_num_split",
	ProjectStartDate:       "project_start_date",
	ProjectTitle:           "project_title",
	SpendingCategories:     "spending_categories",
	SpendingCategoriesDesc: "spending_categories_desc",
	SubprojectId:           "subproject_id",
	Terms:                  "terms",
}
