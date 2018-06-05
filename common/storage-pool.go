package common

import (
	"net"
	"fmt"
	"time"
	"bisale/foundation/thrift/pool"
	"bisale/bisale-console-api/thrift/storage"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func openStorageServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open storage service connection error: %s", err.Error()))
		return nil, err
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	standardClient := thrift.NewTStandardClient(iprot, oprot)

	client := storage.NewStorageClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeStroageServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close storage service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close storage client success")
	return nil
}

func GetStorageServiceClient() (s *storage.StorageClient, c *thriftPool.IdleClient) {

	client, err := StorageServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get storage service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Storage client has closed"))
	}

	//err = StorageServicePool.Put(client)
	//
	//if err != nil {
	//	Log.Error(fmt.Printf("Put storage client to pool error: %s", err.Error()))
	//	return
	//}

	Log.Info("Get storage client from pool success")

	return client.Client.(*storage.StorageClient), client
}
