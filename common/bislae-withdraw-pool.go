package common

import (
	"net"
	"fmt"
	"time"
	"bisale/foundation/thrift/pool"
	"bisale/bisale-console-api/thrift/finance"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func openBisaleWithdrawServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	//protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "depositWithdraw")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale withdraw service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := finance.NewTDepositWithdrawServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleWithdrawServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale withdraw service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale withdraw client success")
	return nil
}

func GetBisaleWithdrawServiceClient() (s *finance.TDepositWithdrawServiceClient, c *thriftPool.IdleClient) {

	client, err := BisaleWithdrawServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale withdraw service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale withdraw client has closed"))
	}

	//err = BisaleWithdrawServicePool.Put(client)
	//
	//if err != nil {
	//	Log.Error(fmt.Printf("Put bisale withdraw client to pool error: %s", err.Error()))
	//	return
	//}

	Log.Info("Get bisale withdraw client from pool success")

	return client.Client.(*finance.TDepositWithdrawServiceClient), client
}
