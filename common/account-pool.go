package common

import (
	"net"
	"fmt"
	"time"
	"bisale/foundation/thrift/pool"
	"bisale/bisale-console-api/thrift/account"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func openAccountServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open account service connection error: %s", err.Error()))
		return nil, err
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	standardClient := thrift.NewTStandardClient(iprot, oprot)

	client := account.NewAccountClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeAccountServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close account service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close account client success")
	return nil
}

func GetAccountServiceClient() (c *account.AccountClient) {

	client, err := AccountServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get account service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Account client has closed"))
	}

	err = AccountServicePool.Put(client)

	if err != nil {
		Log.Error(fmt.Printf("Put account client to pool error: %s", err.Error()))
		return
	}

	Log.Info("Get account client from pool success")

	return client.Client.(*account.AccountClient)
}
