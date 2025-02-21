package scraper

import (
	"time"
)

type SearchCriteria struct {
	FromAirport string
	ToAirport   string
	DateFrom    time.Time
	DateTo      time.Time
	Adults      int
	IsReturn    bool
}

type Vendor interface {
	GenerateURLs(criteria SearchCriteria) []string

	Parse(content string) ([]Flight, error)
}

type Flight struct {
	Vendor        string
	FromAirport   string
	ToAirport     string
	DepartureTime time.Time
	ReturnTime    *time.Time
	Price         float64
	Currency      string
	URL           string
}
