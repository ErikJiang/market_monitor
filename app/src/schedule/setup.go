package schedule

import (
	"github.com/robfig/cron"
	"github.com/rs/zerolog/log"
)

func Setup() {
	log.Info().Msg("enter schedule setup")

	c := cron.New()
	c.AddFunc("0 */1 * * * *", func() {
		log.Debug().Msg("Run Task One ...")
		Task1MarketTicker()
	})
	//c.AddFunc("*/20 * * * * *", func() {
	//	log.Debug().Msg("Run Task Two ...")
	//
	//})

	c.Start()
}
