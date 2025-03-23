package types

type Fetcher interface {
	GenerateURL(criteria SearchCriteria) string

	Fetch(url string, outCh *chan string, errCh *chan error)

	Parse(data string, outCh *chan Flight, errCh *chan error)
}
