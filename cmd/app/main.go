package main

import (
	"context"
	"disputes/internal/pkg/app"
)

func main() {
	ctx := context.Background()
	err := app.Run(ctx)
	if err != nil {
		panic(err)
	}
}
