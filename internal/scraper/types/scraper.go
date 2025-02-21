package types

type Scraper interface {
	GenerateURLs(criteria SearchCriteria) []string

	Fetch(url string) string

	Parse(content string) ([]Flight, error)
}
