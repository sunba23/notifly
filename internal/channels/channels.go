package channels

import (
	"github.com/sunba23/notifly/internal/fetcher/types"
	"sync"
)

type Channels struct {
	FetchParseCh chan string
	ParseWriteCh chan types.Flight
	ErrCh        chan error
	once         sync.Once
}

var (
	instance *Channels
	once     sync.Once
)

func GetChannels() *Channels {
	once.Do(func() {
		instance = &Channels{
			FetchParseCh: make(chan string, 10),
			ParseWriteCh: make(chan types.Flight, 10),
			ErrCh:        make(chan error, 10),
		}
	})
	return instance
}

func (fc *Channels) Close() {
	fc.once.Do(func() {
		close(fc.FetchParseCh)
		close(fc.ParseWriteCh)
		close(fc.ErrCh)
	})
}
