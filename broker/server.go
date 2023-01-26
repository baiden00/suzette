package broker

import (
	"context"
	"github.com/baiden00/suzette/broker/gen"
)

type Server struct {
	gen.UnimplementedBrokerServer
	Broker  *Broker
	Address string
}

func (s *Server) HealthCheck(ctx context.Context, in *gen.HealthCheckRequest) (out *gen.HealthCheckResponse, err error) {
	println(in.Status)
	return &gen.HealthCheckResponse{Status: "Success"}, nil
}

func (s *Server) Publish(ctx context.Context, in *gen.PublishRequest) (out *gen.PublishResponse, err error) {

	return &gen.PublishResponse{}, nil
}

func (s *Server) CreateTopic(ctx context.Context, in *gen.CreateTopicRequest) (out *gen.CreateTopicResponse, err error) {

	return &gen.CreateTopicResponse{}, nil
}

func (s *Server) ClearTopic(ctx context.Context, in *gen.ClearTopicRequest) (out *gen.ClearTopicResponse, err error) {

	return &gen.ClearTopicResponse{}, nil
}

func (s *Server) BatchPublish(stream gen.Broker_BatchPublishServer) (err error) {

	return nil
}
