package scheduler

import (
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
)

func SetupScheduler(c *core.Core) {
	lc, _ := time.LoadLocation("Asia/Jakarta")

	s, err := gocron.NewScheduler(gocron.WithLocation(lc))
	if err != nil {
		log.Panic(err)
	}

	s.Start()
}
