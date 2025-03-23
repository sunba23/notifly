package types

import "time"

type SearchCriteria struct {
	FromAirport string
	ToAirport   string
	DateFrom    time.Time
	DateTo      time.Time
	MinDays     int
	MaxDays     int
	Adults      int
	IsReturn    bool
	NotiPrice   float64
}

type Flight struct {
	Vendor             string
	FromIata           string
	ThereDepartureTime time.Time
	ToIata             string
	BackDepartureTime  *time.Time
	Price              float64
}

type RyanairAPIResponse struct {
	Fares []RFare `json:"fares"`
	Size  int     `json:"size"`
}

type RFare struct {
	Outbound RFlight  `json:"outbound"`
	Inbound  *RFlight `json:"inbound"`
	Summary  RSummary `json:"summary"`
}

type RFlight struct {
	DepartureAirport RAirport `json:"departureAirport"`
	ArrivalAirport   RAirport `json:"arrivalAirport"`
	DepartureDate    string   `json:"departureDate"`
	ArrivalDate      string   `json:"arrivalDate"`
	Price            RPrice   `json:"price"`
	FlightKey        string   `json:"flightKey"`
	FlightNumber     string   `json:"flightNumber"`
	PreviousPrice    *RPrice  `json:"previousPrice"`
	PriceUpdated     float64  `json:"priceUpdated"`
}

type RAirport struct {
	CountryName string `json:"countryName"`
	IataCode    string `json:"iataCode"`
	Name        string `json:"name"`
	SeoName     string `json:"seoName"`
	City        RCity  `json:"city"`
}

type RCity struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	CountryCode string `json:"countryCode"`
	MacCode     string `json:"macCode,omitempty"`
}

type RPrice struct {
	Value               float64 `json:"value"`
	ValueMainUnit       string  `json:"valueMainUnit"`
	ValueFractionalUnit string  `json:"valueFractionalUnit"`
	CurrencyCode        string  `json:"currencyCode"`
	CurrencySymbol      string  `json:"currencySymbol"`
}

type RSummary struct {
	Price            RPrice  `json:"price"`
	PreviousPrice    *RPrice `json:"previousPrice"`
	NewRoute         bool    `json:"newRoute"`
	TripDurationDays int     `json:"tripDurationDays"`
}
