package types

type Fetcher interface {
	GenerateURL(criteria SearchCriteria) string

	Fetch(url string, ch chan string)

	Parse(inCh chan string, outCh chan Flight) ([]Flight, error)
}
