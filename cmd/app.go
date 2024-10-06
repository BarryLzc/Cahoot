package cmd

import (
	"context"
	"os"
)

func Run() {
	ctx := context.Background()
	env := os.Getenv("AQM_ENV")
	if len(env) == 0 {
		env = "test"
	}

	// db
	InitDb()
	// web
	InitWeb(ctx)
}
