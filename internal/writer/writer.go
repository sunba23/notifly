package writer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/sunba23/notifly/internal/channels"
	"github.com/sunba23/notifly/internal/fetcher/types"
)

type Writer struct {
	Filepath       string
	BatchSize      int
	Timeout        time.Duration
	SearchCriteria types.SearchCriteria
	Chans          *channels.Channels
	Wg             *sync.WaitGroup
}

func (w *Writer) check(err error) {
	if err != nil {
		w.Chans.ErrCh <- fmt.Errorf("writer error: %v", err)
	}
}

func (w *Writer) GenerateFilepath() {
	bytes, err := json.Marshal(w.SearchCriteria)
	w.check(err)

	basedir, err := os.UserHomeDir()
	w.check(err)
	dirPath := filepath.Join(basedir, ".config", "notifly")
	err = os.MkdirAll(dirPath, 0755)
	w.check(err)
	w.Filepath = filepath.Join(dirPath, string(bytes))
}

func (w *Writer) checkNoti(flight types.Flight) {
	if flight.Price < w.SearchCriteria.NotiPrice {
		// TODO run notifier here
		fmt.Printf("Found flight with matching criteria: %v\n", flight)
	}
}

func (w *Writer) saveToFile(flights []types.Flight) {
	file, err := os.OpenFile(w.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w.check(err)
	defer file.Close()
	flightsJson, err := json.Marshal(flights)
	w.check(err)
	_, err = file.Write(append(flightsJson, '\n'))
	w.check(err)
}

func (w *Writer) Run(ctx context.Context) {
	w.GenerateFilepath()

	w.Wg.Add(1)
	go func(w *Writer) {
		defer w.Wg.Done()

		var flights []types.Flight
		ticker := time.NewTicker(w.Timeout)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				if len(flights) > 0 {
					w.saveToFile(flights)
				}
				return
			case flight, ok := <-w.Chans.ParseWriteCh:
				if !ok {
					if len(flights) > 0 {
						w.saveToFile(flights)
					}
					return
				}

				w.checkNoti(flight)
				flights = append(flights, flight)

				if len(flights) >= w.BatchSize {
					// save when batch reached
					w.saveToFile(flights)
					flights = nil
					ticker.Reset(w.Timeout)
				}

			case <-ticker.C:
				if len(flights) > 0 {
					// save when timeout reached
					w.saveToFile(flights)
					flights = nil
				}
			}
		}
	}(w)
}
