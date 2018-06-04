package common

import (
	"net"
	"fmt"
	"time"
	"bisale/foundation/thrift/pool"
	"bisale/bisale-console-api/thrift/business"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func openBisaleBusinessServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	//protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "businessKyc")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale business service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := business.NewTReformationActivityServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleBusinessServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale business service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale business client success")
	return nil
}

func GetBisaleBusinessServiceClient() (c *business.TReformationActivityServiceClient) {

	client, err := BisaleBusinessServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale business service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale business client has closed"))
	}

	err = BisaleBusinessServicePool.Put(client)

	if err != nil {
		Log.Error(fmt.Printf("Put bisale business client to pool error: %s", err.Error()))
		return
	}

	Log.Info("Get bisale business client from pool success")

	return client.Client.(*business.TReformationActivityServiceClient)
}
