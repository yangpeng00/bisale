package common

import (
	"net"
	"fmt"
	"time"
	"bisale/foundation/thrift/pool"
	"git.apache.org/thrift.git/lib/go/thrift"
	"bisale/bisale-console-api/thrift/content"
)

func openBisaleContentServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	//protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "content")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale content service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := content.NewTContentServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleContentServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale content service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale content client success")
	return nil
}

func GetBisaleContentServiceClient() (s *content.TContentServiceClient, c *thriftPool.IdleClient) {

	client, err := BisaleContentServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale content service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale content client has closed"))
	}

	Log.Info("Get bisale content client from pool success")

	return client.Client.(*content.TContentServiceClient), client
}
