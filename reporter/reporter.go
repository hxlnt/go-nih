package reporter

const (
	projectURL     = "https://api.reporter.nih.gov/v2/projects/search"
	publicationURL = "https://api.reporter.nih.gov/v2/publications/search"
)

// type Client struct {
// 	client               *http.Client
// 	projectSearchURL     *url.URL
// 	publicationSearchURL *url.URL
// 	ProjectSearch        *ProjectSearchService
// 	PublicationSearch    *PublicationSearchService
// }

// func NewClient() *Client {
// 	httpClient := &http.Client{}
// 	httpClient2 := *httpClient
// 	c := &Client{client: &httpClient2}
// 	c.initialize()
// 	return c
// }
