package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/sunba23/notifly/internal/fetcher"
	"github.com/sunba23/notifly/internal/fetcher/types"
	// "github.com/sunba23/notifly/internal/consumer"
	// "github.com/sunba23/notifly/internal/notifier"
)

var (
	fromAirport string
	toAirport   string
	dateFrom    string
	dateTo      string
	isReturn    bool
	adults      int
	email        string
	notifyPrice int
	minDays     int
	maxDays     int
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Start monitoring specified flights",
	Run:   run,
}

func init() {
	monitorCmd.Flags().StringVar(&fromAirport, "from", "", "Origin airport (e.g., WRO)")
	monitorCmd.Flags().StringVar(&toAirport, "to", "", "Destination airport (e.g., STN)")
	monitorCmd.Flags().StringVar(&dateFrom, "date-from", "", "Earliest date (YYYY-MM-DD)")
	monitorCmd.Flags().StringVar(&dateTo, "date-to", "", "Latest date (YYYY-MM-DD)")
	monitorCmd.Flags().IntVar(&adults, "adults", 1, "Number of adults (e.g., 2)")

	monitorCmd.Flags().BoolVar(&isReturn, "return", true, "Is return (default: true)")
	monitorCmd.Flags().IntVar(&minDays, "min-days", 2, "Min. amount of days")
	monitorCmd.Flags().IntVar(&maxDays, "max-days", 7, "Max. amount of days")

	monitorCmd.Flags().IntVar(&notifyPrice, "noti-price", 0, "Price, below which you will get notified")
	monitorCmd.Flags().StringVar(&email, "email", "", "email to send notifications to")

	monitorCmd.MarkFlagRequired("from")
	monitorCmd.MarkFlagRequired("to")
	monitorCmd.MarkFlagRequired("date-from")
	monitorCmd.MarkFlagRequired("date-to")

	monitorCmd.MarkFlagRequired("noti-price")
	monitorCmd.MarkFlagRequired("email")

	rootCmd.AddCommand(monitorCmd)
}

func parseArguments() (types.SearchCriteria, error) {
	dateFromParsed, err := time.Parse("2006-01-02", dateFrom)
	if err != nil {
		return types.SearchCriteria{}, fmt.Errorf("could not parse \"%v\" to format YYYY-MM-DD: %w", dateFrom, err)
	}

	dateToParsed, err := time.Parse("2006-01-02", dateTo)
	dateToParsed = dateToParsed.Add(24*time.Hour - time.Nanosecond)
	if err != nil {
		return types.SearchCriteria{}, fmt.Errorf("could not parse \"%v\" to format YYYY-MM-DD: %w", dateTo, err)
	}

	config := types.SearchCriteria{
		FromAirport: fromAirport,
		ToAirport:   toAirport,
		DateFrom:    dateFromParsed,
		DateTo:      dateToParsed,
		IsReturn:    isReturn,
		Adults:      adults,
	}

	return config, nil
}

func run(cmd *cobra.Command, args []string) {
	//TODO: run fetcher (fetch, parse), writer (save to disk), notifier goroutines

	searchCriteria, err := parseArguments()

	if err != nil {
		return
	}

	fetcher.Run(searchCriteria)
}
