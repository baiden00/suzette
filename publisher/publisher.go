package publisher

import (
	"github.com/baiden00/suzette/broker/gen"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// TODO: CHANGE ALL ERROR STRINGS TO ACTUAL ERRORS WITH CONSTANTS

const (
	ErrorNotOk = "NO ERROR"
	ErrorOk    = "ERROR"
)

type Publisher struct {
	mu                       sync.Mutex
	id                       string
	targetAddress            string
	connOpts                 grpc.DialOption
	errorHandler             CustomError
	brokerGrpcClient         gen.BrokerClient
	closeConnSignalFromError chan bool
	closeConnSignalFromUser  chan os.Signal
	clientConnection         *grpc.ClientConn
}

type CustomError struct{}

func (m *CustomError) ErrorNotOK() string {
	return ErrorOk //Create Error Codes
}
func (m *CustomError) ErrorOK() string {
	return ErrorNotOk //Create Error Codes
}

func (p *Publisher) NewPublisher(targetAddress string) (*Publisher, string) {
	publisher := &Publisher{
		id:                       generateGUID(),
		connOpts:                 grpc.WithTransportCredentials(insecure.NewCredentials()),
		targetAddress:            targetAddress,
		closeConnSignalFromError: make(chan bool),
		closeConnSignalFromUser:  make(chan os.Signal, 1),
	}

	// Create the GRPC Client
	client, conn, err := p.GenerateBrokerClient()
	if err != ErrorOk {
		return nil, ErrorNotOk
	}
	p.brokerGrpcClient, p.clientConnection = client, conn

	//Setup Notification mechanism for Keyboard Interrupts
	signal.Notify(p.closeConnSignalFromUser, syscall.SIGINT, syscall.SIGTERM)

	return publisher, ErrorOk

}
func (p *Publisher) Publish() {

}

func (p *Publisher) CreateTopic() {

}

func (p *Publisher) ClearTopic() {

}

func (p *Publisher) BatchPublish() {

}

func generateGUID() string {
	generatedUUID := uuid.New()
	return generatedUUID.String()

}
func (p *Publisher) GenerateBrokerClient() (gen.BrokerClient, *grpc.ClientConn, string) {
	// If publisher is not initialized don't do anything else it will panic.
	if p != nil {
		conn, err := grpc.Dial(p.targetAddress, p.connOpts)
		if err != nil {
			return nil, nil, p.errorHandler.ErrorNotOK()
		} else {
			client := gen.NewBrokerClient(conn)
			return client, conn, p.errorHandler.ErrorOK()
		}
	}
	return nil, nil, ErrorNotOk

}

// TODO Revisit, rethink and reimplement. Too tired.
//func (p *Publisher) CloseConnection() {
//	go func() {
//		select {
//		case <-p.closeConnSignalFromError:
//			err := p.clientConnection.Close()
//			if err != nil {
//				os.Exit(1)
//				return
//			}
//		case <-p.closeConnSignalFromUser:
//			err := p.clientConnection.Close()
//			if err != nil {
//				//cleanup resources
//				os.Exit(1)
//				return
//			}
//		}
//	}()
//}
