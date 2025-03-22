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

func Run(criteria types.SearchCriteria) {
	fetcherSlice := []types.Fetcher{
		fetchers.NewRyanairFetcher(),
		//fetchers.NewWizzairFetcher(),
	}

	chans := channels.GetChannels()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// runs web fetchers wrapped in goroutines
	for _, f := range fetcherSlice {
		url := f.GenerateURL(criteria)
		wg.Add(1)
		go func(f types.Fetcher) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					go f.Fetch(url, chans.FetchParseCh, chans.ErrCh)
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
			for rawData := range chans.FetchParseCh {
				f.Parse(rawData, chans.ParseWriteCh, chans.ErrCh)
			}
			for {
				select {
				case <-ctx.Done():
					return
				case rawData, ok := <-chans.FetchParseCh:
					if !ok {
						return
					}
					f.Parse(rawData, chans.ParseWriteCh, chans.ErrCh)
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
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case err, ok := <-chans.ErrCh:
				if !ok {
					return
				}
				fmt.Println("Error:", err)
			}
		}
	}()

	wg.Wait()

  chans.Close()
}
