package cmd

import (
	"context"
	"github.com/english-learning/internal/pkg/web"
)

func InitWeb(ctx context.Context) {
	r := web.InitRouter()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
