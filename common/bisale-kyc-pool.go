package common

import (
	"net"
	"fmt"
	"time"
	"bisale/foundation/thrift/pool"
	"bisale/bisale-console-api/thrift/user"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func openBisaleUserKycServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	//protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "userKyc")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale kyc service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := user.NewTUserKycServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleUserKycServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale kyc service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale user client success")
	return nil
}

func GetBisaleUserKycServiceClient() (s *user.TUserKycServiceClient, c *thriftPool.IdleClient) {

	client, err := BisaleUserKycServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale kyc service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale kyc client has closed"))
	}

	Log.Info("Get bisale kyc client from pool success")

	return client.Client.(*user.TUserKycServiceClient), client
}
