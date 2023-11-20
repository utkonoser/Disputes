package app

import (
	"context"
	adminapi "disputes/internal/app/api/admin"
	publicapi "disputes/internal/app/api/public"
	config "disputes/internal/config/app"
	"github.com/utrack/clay/v2/server"
	"google.golang.org/grpc"
	"log"
	"sync"
)

type DisputesApp struct {
	publicAPI *publicapi.Implementation
	adminApi  *adminapi.Implementation

	unaryInterceptor []grpc.UnaryServerInterceptor
}

func Run(ctx context.Context) error {

	cfg := config.New()

	disputesApp, err := newApp(ctx, cfg)
	if err != nil {
		return err
	}

	return disputesApp.Build()
}

func newApp(ctx context.Context, cfg config.Config) (DisputesApp, error) {
	return DisputesApp{
		publicAPI:        publicapi.NewPublicAPI(&cfg),
		adminApi:         adminapi.NewAdminAPI(&cfg),
		unaryInterceptor: nil,
	}, nil
}

func (d *DisputesApp) Build() (err error) {
	err = d.startAPI()
	if err != nil {
		return err
	}
	return nil
}

func (d *DisputesApp) startAPI() (err error) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func(err error) {
		defer wg.Done()
		adminServer := server.NewServer(8085)
		log.Println("Admin API start!")
		err = adminServer.Run(d.adminApi)
		log.Println("Admin API stop!")

	}(err)

	go func(err error) {
		defer wg.Done()
		publicServer := server.NewServer(8084)
		log.Println("Public API start!")
		err = publicServer.Run(d.publicAPI)
		log.Println("Public API stop!")
	}(err)
	wg.Wait()
	return err
}
