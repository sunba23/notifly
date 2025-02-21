package scrapers

import (
	"fmt"
	"github.com/sunba23/notifly/internal/scraper/types"
	"net/url"
)

type RyanairScraper struct {
	baseURL string
}

func NewRyanairScraper() types.Scraper {
	return &RyanairScraper{
		baseURL: "https://www.ryanair.com/pl/pl/fare-finder",
	}
}

func (r *RyanairScraper) GenerateURLs(criteria types.SearchCriteria) []string {
	params := url.Values{}
	params.Add("originIata", criteria.FromAirport)
	params.Add("destinationIata", criteria.ToAirport)
	params.Add("isReturn", fmt.Sprintf("%t", criteria.IsReturn))
	params.Add("adults", fmt.Sprintf("%d", criteria.Adults))
	params.Add("dateOut", criteria.DateFrom.Format("2006-01-02"))
	params.Add("dateIn", criteria.DateTo.Format("2006-01-02"))

	return []string{fmt.Sprintf("%s?%s", r.baseURL, params.Encode())}
}

func (r *RyanairScraper) Fetch(url string) string {
  return ""
}

func (r *RyanairScraper) Parse(content string) ([]types.Flight, error) {
	return nil, nil
}
