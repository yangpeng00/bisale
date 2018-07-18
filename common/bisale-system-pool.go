package common

import (
	"time"
	"bisale/foundation/thrift/pool"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"bisale/bisale-console-api/thrift/system"
)

func openBisaleSystemServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	//protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "system")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale system service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := system.NewTSystemServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleSystemServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale system service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale user client success")
	return nil
}

func GetBisaleSystemServiceClient() (s *system.TSystemServiceClient, c *thriftPool.IdleClient) {

	client, err := BisaleSystemServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale system service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale system client has closed"))
	}

	Log.Info("Get bisale system client from pool success")

	return client.Client.(*system.TSystemServiceClient), client
}


