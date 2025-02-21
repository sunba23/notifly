package scraper

import (
	"fmt"

	"github.com/sunba23/notifly/internal/scraper/scrapers"
	"github.com/sunba23/notifly/internal/scraper/types"
)

func RunScrapers(criteria types.SearchCriteria) {
	scraperSlice := []types.Scraper{
		scrapers.NewRyanairScraper(),
		scrapers.NewWizzairScraper(),
	}

  var urls []string
	for _, s := range scraperSlice {
    urls = s.GenerateURLs(criteria)
	}
  fmt.Println(urls)

	// run scrapers goroutines - fetch, parse, publish
}
