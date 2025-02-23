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

func (r *WizzairFetcher) Fetch(url string, ch chan string) string {
  return ""
}

func (w *WizzairFetcher) Parse(ch chan string) ([]types.Flight, error) {
	return nil, nil
}
