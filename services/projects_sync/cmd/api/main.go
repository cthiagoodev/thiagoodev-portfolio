package main

import (
	"context"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/module"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	syncModule, err := module.NewSyncModule(ctx)
	if err != nil {
		panic(err)
	}

	defer func(m *module.SyncModule) {
		m.DbPool.Close()
		m.SbPool.Close()
	}(syncModule)

	err = syncModule.UseCase.Execute(ctx)
	if err != nil {
		panic(err)
	}
}
