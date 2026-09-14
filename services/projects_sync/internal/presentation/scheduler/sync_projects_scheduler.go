package scheduler

import (
	"context"
	"log"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/usecases"
	"github.com/robfig/cron/v3"
)

type SyncProjectsScheduler struct {
	syncProjectsUseCase usecases.SyncProjectsUseCase
}

func NewSyncProjectsScheduler(syncProjectsUseCase usecases.SyncProjectsUseCase) *SyncProjectsScheduler {
	return &SyncProjectsScheduler{syncProjectsUseCase}
}

func (s *SyncProjectsScheduler) Schedule() {
	ctx := context.Background()

	c := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cron.DefaultLogger),
	))

	id, err := c.AddFunc("0 3 * * *", func() {
		err := s.syncProjectsUseCase.Execute(ctx)
		if err != nil {
			log.Printf("Failed to execute sync projects for projects: %v", err)
			return
		}
	})

	if err != nil {
		c.Stop()
		c.Remove(id)
		return
	}

	c.Start()
}
