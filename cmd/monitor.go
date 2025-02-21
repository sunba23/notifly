package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/sunba23/notifly/internal/scraper"
	// "github.com/sunba23/notifly/internal/consumer"
	// "github.com/sunba23/notifly/internal/notifier"
)

var (
	fromAirport string
	toAirport   string
	dateFrom    string
	dateTo      string
	notifyType  string
	adults      int
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Start monitoring specified flights",
	Run:   run,
}

func init() {
	monitorCmd.Flags().StringVar(&fromAirport, "from", "", "Origin airport (e.g., WRO)")
	monitorCmd.Flags().StringVar(&toAirport, "to", "", "Destination airport (e.g., MRS)")
	monitorCmd.Flags().StringVar(&dateFrom, "date-from", "", "Start date (YYYY-MM-DD)")
	monitorCmd.Flags().StringVar(&dateTo, "date-to", "", "End date (YYYY-MM-DD)")
	monitorCmd.Flags().StringVar(&notifyType, "notify", "file", "Notification method (email, file)")
	monitorCmd.Flags().IntVar(&adults, "adults", 1, "Number of adults (e.g., 2)")

	monitorCmd.MarkFlagRequired("from")
	monitorCmd.MarkFlagRequired("to")
	monitorCmd.MarkFlagRequired("date-from")
	monitorCmd.MarkFlagRequired("date-to")

	rootCmd.AddCommand(monitorCmd)
}

func parseArguments() (scraper.SearchCriteria, error) {
	dateFromParsed, err := time.Parse("2006-01-02", dateFrom)
	if err != nil {
		return scraper.SearchCriteria{}, fmt.Errorf("could not parse \"%v\" to format YYYY-MM-DD: %w", dateFrom, err)
	}

	dateToParsed, err := time.Parse("2006-01-02", dateTo)
  dateToParsed = dateToParsed.Add(24 * time.Hour - time.Nanosecond)
	if err != nil {
		return scraper.SearchCriteria{}, fmt.Errorf("could not parse \"%v\" to format YYYY-MM-DD: %w", dateTo, err)
	}

	config := scraper.SearchCriteria{
		FromAirport: fromAirport,
		ToAirport:   toAirport,
		DateFrom:    dateFromParsed,
		DateTo:      dateToParsed,
		Adults:      adults,
	}

	return config, nil
}

func run(cmd *cobra.Command, args []string) {
  //TODO: run scraper (fetch, parse, publish), consumer (subscribe, process, saveDb) and notifier (readDb, notify) goroutines

	scraperConfig, err := parseArguments()

  if err != nil {
    return
  }

	scraper.RunScraper(scraperConfig)
}
