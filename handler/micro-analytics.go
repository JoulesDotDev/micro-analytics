package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	microanalytics "micro-analytics/proto"
)

type MicroAnalytics struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *MicroAnalytics) Call(ctx context.Context, req *microanalytics.Request, rsp *microanalytics.Response) error {
	log.Info("Received MicroAnalytics.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *MicroAnalytics) Stream(ctx context.Context, req *microanalytics.StreamingRequest, stream microanalytics.MicroAnalytics_StreamStream) error {
	log.Infof("Received MicroAnalytics.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&microanalytics.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *MicroAnalytics) PingPong(ctx context.Context, stream microanalytics.MicroAnalytics_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&microanalytics.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
