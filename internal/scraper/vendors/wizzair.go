package vendors

import "github.com/sunba23/notifly/internal/scraper"

type WizzairVendor struct {
	baseURL string
}

func NewWizzairVendor() *WizzairVendor {
	return &WizzairVendor{
		baseURL: "https://wizzair.com/search",
	}
}

func (w *WizzairVendor) GenerateURLs(criteria scraper.SearchCriteria) []string {
	return nil
}

func (w *WizzairVendor) Parse(content string) ([]scraper.Flight, error) {
	return nil, nil
}
