package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/m3o/m3o/services/wallet/handler"
	pb "github.com/m3o/m3o/services/wallet/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("wallet"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterWalletHandler(srv.Server(), handler.NewHandler(srv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
