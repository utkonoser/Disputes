package app

import (
	"context"
	adminapi "disputes/internal/app/api/admin"
	publicapi "disputes/internal/app/api/public"
	config "disputes/internal/config/app"
	"google.golang.org/grpc"
)

var (
	Name    = "disputes"
	Version = "DEV"
)

type DisputesApp struct {
	publicAPI *publicapi.Implementation
	adminApi  *adminapi.Implementation

	unaryInterceptor []grpc.UnaryServerInterceptor
}

func Run(ctx context.Context) error {

	cfg := config.New()

	disputesApp := newApp(ctx, cfg)

	return disputesApp.run()
}

func newApp(ctx context.Context, cfg config.Config) DisputesApp {
	return DisputesApp{
		publicAPI:        nil,
		adminApi:         nil,
		unaryInterceptor: nil,
	}
}

func (d *DisputesApp) run() error {
	// not implemented
	return nil
}
