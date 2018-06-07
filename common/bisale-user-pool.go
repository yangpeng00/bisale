package common

import (
	"bisale/foundation/thrift/pool"
	"net"
	"fmt"
	"time"
	"git.apache.org/thrift.git/lib/go/thrift"
	"bisale/bisale-console-api/thrift/user"
)

func openBisaleUserServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	//protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "user")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale user service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := user.NewTUserServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleUserServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale user service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale user client success")
	return nil
}

func GetBisaleUserServiceClient() (s *user.TUserServiceClient, c *thriftPool.IdleClient) {

	client, err := BisaleUserServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale user service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale user client has closed"))
	}

	//err = BisaleUserServicePool.Put(client)
	//
	//if err != nil {
	//	Log.Error(fmt.Printf("Put bisale user client to pool error: %s", err.Error()))
	//	return
	//}
	//
	Log.Info("Get bisale user client from pool success")

	return client.Client.(*user.TUserServiceClient), client
}
