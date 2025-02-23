package fetchers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/sunba23/notifly/internal/fetcher/types"
)

type RyanairFetcher struct {
	apiURL string
}

type RyanairAPIResponse struct {
	Fares []struct {
		Price struct {
			Value float64 `json:"value"`
		} `json:"price"`
		DepartureDate string `json:"departureDate"`
	} `json:"fares"`
}

func NewRyanairFetcher() types.Fetcher {
	return &RyanairFetcher{
		apiURL: "https://www.ryanair.com/api/farfnd/v4/roundTripFares",
	}
}

func (r *RyanairFetcher) GenerateURL(criteria types.SearchCriteria) string {
	params := url.Values{}
	params.Set("departureAirportIataCode", criteria.FromAirport)
  // earliest to go there = dateFrom
	params.Set("outboundDepartureDateFrom", criteria.DateFrom.Format("2006-01-02"))
	params.Set("market", "pl-pl")
	params.Set("adultPaxCount", fmt.Sprintf("%d", criteria.Adults))
	params.Set("arrivalAirportIataCodes", criteria.ToAirport)
  params.Set("searchMode", "ALL")
  // latest to go there = dateTo - minimum days spent
	params.Set("outboundDepartureDateTo", criteria.DateTo.AddDate(0,0, -1 * criteria.MinDays).Format("2006-01-02"))
  // earliest to go back = dateFrom + min days spent
	params.Set("inboundDepartureDateFrom", criteria.DateFrom.AddDate(0, 0, 1).Format("2006-01-02"))
  // latest to go back = dateTo
	params.Set("inboundDepartureDateTo", criteria.DateTo.Format("2006-01-02"))
  params.Set("durationFrom", strconv.Itoa(criteria.MinDays))
  params.Set("durationTo", strconv.Itoa(criteria.MaxDays))
  params.Set("outboundDepartureDaysOfWeek", "MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY,SUNDAY")
  params.Set("outboundDepartureTimeFrom", "00:00")
  params.Set("outboundDepartureTimeTo", "23:59")
  params.Set("inboundDepartureTimeFrom", "00:00")
  params.Set("inboundDepartureTimeTo", "23:59")

	return fmt.Sprintf("%s?%s", r.apiURL, params.Encode())
}

func (r *RyanairFetcher) Fetch(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
  ch <- string(body)
}

func (r *RyanairFetcher) Parse(ch chan string) ([]types.Flight, error) {
	var response RyanairAPIResponse
	if err := json.Unmarshal([]byte(<-ch), &response); err != nil {
		return nil, err
	}

	var flights []types.Flight
	for _, fare := range response.Fares {
		departureTime, _ := time.Parse(time.RFC3339, fare.DepartureDate)
		flights = append(flights, types.Flight{
			Vendor:        "Ryanair",
			FromAirport:   criteria.FromAirport,
			ToAirport:     criteria.ToAirport,
			Price:         fare.Price.Value,
			DepartureTime: departureTime,
			URL:           "https://www.ryanair.com/",
		})
	}
	return flights, nil
}
