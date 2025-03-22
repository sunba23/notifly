package fetcher

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sunba23/notifly/internal/fetcher/fetchers"
	"github.com/sunba23/notifly/internal/fetcher/types"
)

func Run(criteria types.SearchCriteria) {
	fetcherSlice := []types.Fetcher{
		fetchers.NewRyanairFetcher(),
		fetchers.NewWizzairFetcher(),
	}

	fetchParseCh := make(chan string, 10)
	parsePublishCh := make(chan types.Flight, 10)
	errCh := make(chan error, 10)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// runs web fetchers wrapped in goroutines
	for _, f := range fetcherSlice {
		wg.Add(1)
		go func(f types.Fetcher) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					url := f.GenerateURL(criteria)
					go f.Fetch(url, fetchParseCh, errCh)
					time.Sleep(1 * time.Minute)
				}
			}
		}(f)
	}

	// runs parsers wrapped in goroutines
	for _, f := range fetcherSlice {
		wg.Add(1)
		go func(f types.Fetcher) {
			defer wg.Done()
			// parses data, simultaneously pushing each parsed flight into next channel
			for rawData := range fetchParseCh {
				f.Parse(rawData, parsePublishCh, errCh)
			}
			for {
				select {
				case <-ctx.Done():
					return
				case rawData, ok := <-fetchParseCh:
					if !ok {
						return
					}
					f.Parse(rawData, parsePublishCh, errCh)
				}
			}
		}(f)
	}

	// ONLY FOR TESTS - print out parsed flights
	wg.Add(1)
	go func() {
		defer wg.Done()
		var incr uint8 = 0
		for {
			select {
			case <-ctx.Done():
				return
			case parsedFlight, ok := <-parsePublishCh:
				if !ok {
					return
				}
				fmt.Printf("parsed flight %v: %v\n", incr, parsedFlight)
				incr++
			}
		}
	}()

	// run publishing goroutine
	// wg.Add(1)
	// go func (){
	//   for flight := range parsePublishCh {
	//   }
	// }()

	// runs errors goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case err, ok := <-errCh:
				if !ok {
					return
				}
				fmt.Println("Error:", err)
			}
		}
	}()

	wg.Wait()

	close(fetchParseCh)
	close(parsePublishCh)
	close(errCh)
}
