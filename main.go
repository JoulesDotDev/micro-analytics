package main

import (
	"micro-analytics/handler"
	pb "micro-analytics/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("micro-analytics"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterMicroAnalyticsHandler(srv.Server(), new(handler.MicroAnalytics))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
