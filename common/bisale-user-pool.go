package common

import (
	"net"
	"fmt"
	"time"
	"bisale/foundation/thrift/pool"
	"bisale/bisale-console-api/thrift/user"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func openBisaleServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {


	//protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp :=thrift.NewTMultiplexedProtocol(protocol,"userKyc")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale user service connection error: %s", err.Error()))
	}

	// iprot := protocolFactory.GetProtocol(transport)
	// oprot := protocol.GetProtocol(transport)

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := user.NewTUserKycServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	fmt.Println(err.Error())
	Log.Error(fmt.Printf("Close bisale user service connection error: %s", err.Error()))
	return err
}

func GetBisaleServiceClient() (c *user.TUserKycServiceClient) {

	client, err := BisaleUserServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale user service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale user client has closed"))
	}

	err = BisaleUserServicePool.Put(client)

	if err != nil {
		Log.Error(fmt.Printf("Put bisale user client to pool error: %s", err.Error()))
		return
	}

	Log.Info("Get bisale user client from pool success")

	return client.Client.(*user.TUserKycServiceClient)
}
