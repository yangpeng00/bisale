package common

import (
	"time"
	"bisale/foundation/thrift/pool"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"bisale/bisale-console-api/thrift/engine"
)

func openBisaleOrderServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	//protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "orders")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale order service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := engine.NewTOrdersServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleOrderServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale order service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale user client success")
	return nil
}

func GetBisaleOrderServiceClient() (s *engine.TOrdersServiceClient, c *thriftPool.IdleClient) {

	client, err := BisaleOrderServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale order service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale order client has closed"))
	}

	Log.Info("Get bisale order client from pool success")

	return client.Client.(*engine.TOrdersServiceClient), client
}

