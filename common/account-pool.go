package common

import (
	"net"
	"fmt"
	"time"
	"bisale/foundation/thrift/pool"
	"bisale/thrift-account/thrift/account"
	"git.apache.org/thrift.git/lib/go/thrift"
	"bisale/thrift-message/thrift/message"
)

func openAccountServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open account service connection error: %s", err))
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
	Log.Error(fmt.Printf("Close account service connection error: %s", err))
	return err
}

func getAccountServiceClient(messageClientPool *thriftPool.ThriftPool) (c *message.MessageClient) {

	client, err := messageClientPool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get account service client error: %s", err))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Account client has closed"))
	}

	err = messageClientPool.Put(client)

	if err != nil {
		Log.Error(fmt.Printf("Put account client to pool error: %s", err))
		return
	}

	return client.Client.(*message.MessageClient)
}
