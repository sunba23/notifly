package scrapers

import "github.com/sunba23/notifly/internal/scraper/types"

type WizzairScraper struct {
	baseURL string
}

func NewWizzairScraper() types.Scraper {
	return &WizzairScraper{
		baseURL: "https://wizzair.com/search",
	}
}

func (w *WizzairScraper) GenerateURLs(criteria types.SearchCriteria) []string {
	return nil
}

func (r *WizzairScraper) Fetch(url string) string {
  return ""
}

func (w *WizzairScraper) Parse(content string) ([]types.Flight, error) {
	return nil, nil
}
