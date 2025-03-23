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
	params.Set("outboundDepartureDateTo", criteria.DateTo.AddDate(0, 0, -1*criteria.MinDays).Format("2006-01-02"))
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

func (r *RyanairFetcher) Fetch(url string, outCh *chan string, errCh *chan error) {
	resp, err := http.Get(url)
	if err != nil {
		*errCh <- fmt.Errorf("Error fetching %s: %v\n", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		*errCh <- fmt.Errorf("Error reading response body: %v\n", err)
	}
	*outCh <- string(body)
}

func (r *RyanairFetcher) Parse(data string, outCh *chan types.Flight, errCh *chan error) {
	var response types.RyanairAPIResponse
	if err := json.Unmarshal([]byte(data), &response); err != nil {
		*errCh <- err
	}

	if response.Size == 0 {
		*errCh <- fmt.Errorf("no flights were found!")
	}

	for _, fare := range response.Fares {
		thereDT, err := time.Parse("2006-01-02T15:04:05", fare.Outbound.DepartureDate)
		if err != nil {
			*errCh <- fmt.Errorf("failed to parse outbound departure date %v \n", thereDT)
		}

		var backDT *time.Time
		if fare.Inbound != nil {
			parsedBackDT, err := time.Parse("2006-01-02T15:04:05", fare.Inbound.DepartureDate)
			if err != nil {
				*errCh <- fmt.Errorf("failed to parse inbound departure date: %v", err)
			}
			backDT = &parsedBackDT
		}

		*outCh <- types.Flight{
			Vendor:             "Ryanair",
			FromIata:           fare.Outbound.DepartureAirport.IataCode,
			ThereDepartureTime: thereDT,
			ToIata:             fare.Outbound.ArrivalAirport.IataCode,
			BackDepartureTime:  backDT,
			Price:              fare.Summary.Price.Value,
		}
	}
}
