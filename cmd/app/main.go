// Package main is the application entry point.
package main

import (
	"context"
	"golang-lua-integration/internal/di"
	"os"
)

func main() {
	ctx := context.Background()
	srv := di.Init(ctx)
	srv.Run(ctx, os.Args)
}
