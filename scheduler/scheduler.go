package scheduler

import (
	"log"
	"time"

	"github.com/ansel1/merry/v2"
	"github.com/go-co-op/gocron"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
)

func SetupScheduler(c *core.Core) {
	lc, _ := time.LoadLocation("Asia/Jakarta")

	s := gocron.NewScheduler(lc)

	_, err := s.Every(1).Minutes().Do(SendHealthCheckSignal, c)
	if err != nil {
		err = merry.Wrap(err)
		log.Panic(err)
	}

	s.StartAsync()
}
