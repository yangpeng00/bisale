package common

import (
	"net"
	"fmt"
	"time"
	"bisale/foundation/thrift/pool"
	"bisale/thrift-message/thrift/message"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func openMessageServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open message service connection error: %s", err.Error()))
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	standardClient := thrift.NewTStandardClient(iprot, oprot)

	client := message.NewMessageClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeMessageServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	Log.Error(fmt.Printf("Close message service connection error: %s", err.Error()))
	return err
}

func GetMessageServiceClient() (c *message.MessageClient) {

	client, err := MessageServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get message service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Account client has closed"))
	}

	err = MessageServicePool.Put(client)

	if err != nil {
		Log.Error(fmt.Printf("Put message client to pool error: %s", err.Error()))
		return
	}

	Log.Info("Get message client from pool success")

	return client.Client.(*message.MessageClient)
}
