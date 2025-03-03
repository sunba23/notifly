package fetcher

import (
	"github.com/sunba23/notifly/internal/fetcher/fetchers"
	"github.com/sunba23/notifly/internal/fetcher/types"
)

func RunFetchers(criteria types.SearchCriteria) {
	fetcherSlice := []types.Fetcher{
		fetchers.NewRyanairFetcher(),
		fetchers.NewWizzairFetcher(),
	}

  fetchParseCh := make(chan string)
  parsePublishCh := make(chan types.Flight, 5)
	// run fetching, parsing and publishing goroutines
	for _, f := range fetcherSlice {
    // this should periodically fetch data from url and push it into channel
		go f.Fetch(f.GenerateURL(criteria), fetchParseCh)
    // this should get data from channel whenever its there, and parse it into FlightData and push into channel
    go f.Parse(fetchParseCh, parsePublishCh)
	}

	// run publishing goroutine
}
