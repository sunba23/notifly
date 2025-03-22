package writer

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/sunba23/notifly/internal/channels"
	"github.com/sunba23/notifly/internal/fetcher/types"
)

type Writer struct {
	searchCriteria types.SearchCriteria
	path           string
	filename       string
	chans          channels.Channels
	flightPrice    map[string]int
}

func (w *Writer) check(err error) {
	if err != nil {
		w.chans.ErrCh <- fmt.Errorf("writer error: %v", err)
	}
}

func (w *Writer) GenerateFilename() {
	bytes, err := json.Marshal(w.searchCriteria)
	w.check(err)
	w.filename = string(bytes)
}

// saves current prices for flight in in-memory dictionary
// tells notifier to notify and shuts down program
func (w *Writer) checkNoti(flight types.Flight) {

}

// saves batch of flights to file
func (w *Writer) save([]types.Flight) {

}

// main runner
func (w *Writer) Run() {
	chans := channels.GetChannels()
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w.GenerateFilename()

	wg.Add(1)
	go func(w *Writer) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case flight, ok := <-chans.ParseWriteCh:
				if !ok {
					return
				}
				w.checkNoti(flight)
        // TODO gather flights from channel here. after limit max, save. after timeout, save.
				w.save(flights)
			}
		}
	}(w)
}
