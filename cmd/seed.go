package cmd

import (
	"flag"
	"payment-service/db/seeds"
	"payment-service/internal/adapter"

	"github.com/rs/zerolog/log"
)

// RunSeed function is used to run the seed
// It will initialize the database connection and run the seed
func RunSeed(cmd *flag.FlagSet, args []string) {
	var (
		table = cmd.String("table", "", "seed to run")          // ex: go run main.go seed -table=product_categories
		total = cmd.Int("total", 1, "total of records to seed") // ex: go run main.go seed -table=product_categories -total=10
	)

	if err := cmd.Parse(args); err != nil { // parse the flags
		log.Fatal().Err(err).Msg("Error while parsing flags")
	}

	adapter.Adapters.Sync(
		adapter.WithShopeefunPaymentPostgres(),
	)
	defer func() {
		if err := adapter.Adapters.Unsync(); err != nil {
			log.Fatal().Err(err).Msg("Error while unsyncing adapters")
		}
	}()

	seeds.Execute(adapter.Adapters.ShopeefunPaymentPostgres, *table, *total)
}
