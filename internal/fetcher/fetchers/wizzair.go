package fetchers

import "github.com/sunba23/notifly/internal/fetcher/types"

type WizzairFetcher struct {
	baseURL string
}

func NewWizzairFetcher() types.Fetcher {
	return &WizzairFetcher{
		baseURL: "https://wizzair.com/search",
	}
}

func (w *WizzairFetcher) GenerateURL(criteria types.SearchCriteria) string {
	return ""
}

func (r *WizzairFetcher) Fetch(url string, outCh *chan string, errCh *chan error) {
	return
}

func (w *WizzairFetcher) Parse(data string, outCh *chan types.Flight, errCh *chan error) {
	return
}
