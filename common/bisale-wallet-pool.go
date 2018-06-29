package common

import (
	"time"
	"bisale/foundation/thrift/pool"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"bisale/bisale-console-api/thrift/wallet"
)

func openWalletServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open wallet service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(protocol, protocol)

	client := wallet.NewWalletClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeWalletServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close wallet service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close wallet client success")
	return nil
}

func GetWalletServiceClient() (s *wallet.WalletClient, c *thriftPool.IdleClient) {

	client, err := WalletServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get wallet service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Wallet client has closed"))
	}

	Log.Info("Get wallet client from pool success")

	return client.Client.(*wallet.WalletClient), client
}

