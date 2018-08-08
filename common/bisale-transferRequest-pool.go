package common

import (
	"time"
	"bisale/foundation/thrift/pool"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"bisale/bisale-console-api/thrift/balanceAccount"
)
func openBisaleTransferRequestServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "transferRequest")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale transfer request service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := balanceAccount.NewTTransferRequestServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleTransferRequestServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale transfer request service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale transfer request client success")
	return nil
}

func GetBisaleTransferRequestServiceClient() (s *balanceAccount.TTransferRequestServiceClient, c *thriftPool.IdleClient) {

	client, err := BisaleTransferRequestServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale transfer request service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale transfer requestFf client has closed"))
	}

	Log.Info("Get bisale business client from pool success")

	return client.Client.(*balanceAccount.TTransferRequestServiceClient), client
}