package broker

import (
	"context"
	"github.com/baiden00/suzette/broker/gen"
	"google.golang.org/grpc"
	"log"
	"net"
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

	//print(in.Message.MessageId)
	//print(in.Message.MessageValue.TypeUrl)
	//
	//msg := &gen2.TestMessage{}
	//
	//err = proto.Unmarshal(in.Message.MessageValue.Value, msg)
	//if err != nil {
	//	println(err)
	//}
	//
	//print(msg.Num)

	//print(msg.ProtoReflect())
	//
	//val, err := NewProtoInstanceFromAny(in.Message.MessageValue)
	//if err != nil {
	//	print(err)
	//}
	////print(cmp.Equal(val, in.Message.MessageValue))
	//print(reflect.TypeOf(val) == reflect.TypeOf(in.Message.MessageValue))
	//print(reflect.TypeOf(in.Message.MessageValue))
	////print(in.Message.MessageValue.TypeUrl)
	//println(val == nil)
	return &gen.PublishResponse{Response: "TEST"}, nil

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

func NewServer() *Server {
	return &Server{}
}

func NewServerWithConfig(opts []interface{}) *Server {
	server := &Server{}
	for _, v := range opts {
		switch v := v.(type) {
		case *Broker:
			server.Broker = v
		case string:
			server.Address = v
		}
	}
	return server
}

func (s *Server) Start() {

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	service := grpc.NewServer()
	gen.RegisterBrokerServer(service, s)

	if err := service.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

//func NewProtoInstanceFromAny(any *anypb.Any) (protoreflect.ProtoMessage, error) {
//	typeURL := any.TypeUrl
//	// Parse the TypeURL to get the package name and message name
//	segments := strings.Split(typeURL, "/")
//	packageName := segments[0]
//	messageName := segments[1]
//	println(messageName)
//	println(packageName)
//	messageType, _ := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(packageName + "." + messageName))
//	if messageType == nil {
//		print("YES")
//		return nil, fmt.Errorf("message %s not found in package %s", messageName, packageName)
//	}
//
//	newValue := messageType.New().Interface()
//	// Unmarshal the value field of the Any proto into the new message instance
//	println("HEY")
//	if err := protojson.Unmarshal(any.Value, newValue); err != nil {
//		return nil, err
//	}
//
//	print("STRING" + newValue.ProtoReflect().Descriptor().FullName())
//
//	return newValue, nil
//}
