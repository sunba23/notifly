package vendors

import (
	"fmt"
	"net/url"
  "github.com/sunba23/notifly/internal/scraper"
)

type RyanairVendor struct {
	baseURL string
}

func NewRyanairVendor() *RyanairVendor {
	return &RyanairVendor{
		baseURL: "https://www.ryanair.com/pl/pl/fare-finder",
	}
}

func (r *RyanairVendor) GenerateURLs(criteria scraper.SearchCriteria) []string {
	params := url.Values{}
	params.Add("originIata", criteria.FromAirport)
	params.Add("destinationIata", criteria.ToAirport)
	params.Add("isReturn", fmt.Sprintf("%t", criteria.IsReturn))
	params.Add("adults", fmt.Sprintf("%d", criteria.Adults))
	params.Add("dateOut", criteria.DateFrom.Format("2006-01-02"))
	params.Add("dateIn", criteria.DateTo.Format("2006-01-02"))
	
	return []string{fmt.Sprintf("%s?%s", r.baseURL, params.Encode())}
}

func (r *RyanairVendor) Parse(content string) ([]scraper.Flight, error) {
	return nil, nil
}
