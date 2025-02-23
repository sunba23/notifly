package types

type Fetcher interface {
	GenerateURL(criteria SearchCriteria) string

	Fetch(url string, ch chan string) string

	Parse(fetchParseCh chan string, parsePublishCh chan Flight) ([]Flight, error)
}
