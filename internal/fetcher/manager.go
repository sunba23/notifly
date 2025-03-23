package fetcher

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sunba23/notifly/internal/channels"
	"github.com/sunba23/notifly/internal/fetcher/fetchers"
	"github.com/sunba23/notifly/internal/fetcher/types"
)

type Fetcher struct {
	Wg *sync.WaitGroup
  Chans *channels.Channels
}

func (fet *Fetcher) Run(ctx context.Context, criteria types.SearchCriteria) {
	fetcherSlice := []types.Fetcher{
		fetchers.NewRyanairFetcher(),
		//fetchers.NewWizzairFetcher(),
	}

	// runs web fetchers wrapped in goroutines
	for _, f := range fetcherSlice {
		url := f.GenerateURL(criteria)
    fmt.Println("generated url")
		fet.Wg.Add(1)
    fmt.Println("added to wg")
		go func(f types.Fetcher) {
			defer fet.Wg.Done()
			for {
				select {
				case <-ctx.Done():
          fmt.Println("ctx done")
					return
				default:
          fmt.Println("fetching flight data")
					go f.Fetch(url, &fet.Chans.FetchParseCh, &fet.Chans.ErrCh)
					time.Sleep(1 * time.Minute)
				}
			}
		}(f)
	}

	// runs parsers wrapped in goroutines
	for _, f := range fetcherSlice {
		fet.Wg.Add(1)
		go func(f types.Fetcher) {
			defer fet.Wg.Done()
			// parses data, simultaneously pushing each parsed flight into next channel
			for rawData := range fet.Chans.FetchParseCh {
				f.Parse(rawData, &fet.Chans.ParseWriteCh, &fet.Chans.ErrCh)
			}
			for {
				select {
				case <-ctx.Done():
					return
				case rawData, ok := <-fet.Chans.FetchParseCh:
					if !ok {
						return
					}
					f.Parse(rawData, &fet.Chans.ParseWriteCh, &fet.Chans.ErrCh)
				}
			}
		}(f)
	}

	// ONLY FOR TESTS - print out parsed flights
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	var incr uint8 = 0
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			return
	// 		case parsedFlight, ok := <-parseWriteCh:
	// 			if !ok {
	// 				return
	// 			}
	// 			fmt.Printf("parsed flight %v: %v\n", incr, parsedFlight)
	// 			incr++
	// 		}
	// 	}
	// }()

	// runs errors goroutine
	fet.Wg.Add(1)
	go func() {
		defer fet.Wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case err, ok := <-fet.Chans.ErrCh:
				if !ok {
					return
				}
				fmt.Println("Error:", err)
			}
		}
	}()
}
