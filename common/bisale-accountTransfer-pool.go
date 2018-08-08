package common

import (
	"time"
	"bisale/foundation/thrift/pool"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"bisale/bisale-console-api/thrift/engine"
)
func openBisaleAccountTransferServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "accountTransfers")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale accountTransfer service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := engine.NewTAccountTransfersServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleAccountTransferServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale accountTransfer service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale accountTransfer client success")
	return nil
}

func GetBisaleAccountTransferServiceClient() (s *engine.TAccountTransfersServiceClient, c *thriftPool.IdleClient) {

	client, err := BisaleAccountTransferServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale accountTransfer service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale accountTransfer client has closed"))
	}

	Log.Info("Get bisale business client from pool success")

	return client.Client.(*engine.TAccountTransfersServiceClient), client
}