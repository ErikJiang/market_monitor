package schedule

import (
	"github.com/rs/zerolog/log"
	"time"
)

var CronManager Manager

type Test struct {
}

func (Test) Frequency() string {
	return "30 * * * * *" // Every 30 seconds
}

func (Test) Run() {
	log.Debug().Msgf("Test Crontab", time.Now().Unix())
}

func Setup() {
	log.Info().Msg("enter schedule setup")
	CronManager := NewManager()
	CronManager.Register(Test{})
	CronManager.RegisterFunc("0 * * * * *", func() { // Register Func
		log.Debug().Msgf("%v", time.Now().Unix())
	})

	CronManager.Start()
	log.Debug().Msg("Cron started success")

}
