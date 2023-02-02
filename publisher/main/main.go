package main

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	//registry := protoregistry.Files{}

	//protoFile, err := ioutil.ReadFile("/Users/charlesbaiden/Github/suzette/publisher/test.proto")
	//if err != nil {
	//	print(err)
	//}
	//
	//pbSet := new(descriptorpb.FileDescriptorSet)
	//if err := proto.Unmarshal(protoFile, pbSet); err != nil {
	//	print(err)
	//}
	//
	//// We know protoc was invoked with a single .proto file.
	//pb := pbSet.GetFile()
	//
	//// Initialize the file descriptor object.
	//fd, err := protodesc.NewFile(pb[0], protoregistry.GlobalFiles)
	//if err != nil {
	//	print(err)
	//}
	//
	//// and finally register it.
	//err = protoregistry.GlobalFiles.RegisterFile(fd)
	//if err != nil {
	//	print(err)
	//}
	//
	//conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//
	//}
	//
	//defer func(conn *grpc.ClientConn) {
	//	err := conn.Close()
	//	if err != nil {
	//		println(err)
	//	}
	//}(conn)
	//
	//c := gen.NewBrokerClient(conn)
	//resp, err := c.HealthCheck(context.Background(), &gen.HealthCheckRequest{
	//	Status: "PING",
	//})
	//
	//if err != nil {
	//	println(err.Error())
	//}
	//println(resp.Status)
	//
	//bytes, err := proto.Marshal(&gen2.TestMessage{MessageId: "xyz", Num: 55})
	//if err != nil {
	//	println(err)
	//}
	//_any := &any2.Any{
	//	TypeUrl: "baiden.xyz/test.proto",
	//	Value:   bytes,
	//}
	//
	//resp2, err := c.Publish(context.Background(), &gen.PublishRequest{
	//	PublisherId: "ama",
	//	Message: &gen.Message{
	//		MessageId:    "key",
	//		MessageValue: _any,
	//	},
	//})
	//
	//if err != nil {
	//	println(err.Error())
	//}
	//println(resp2.Response)

}
