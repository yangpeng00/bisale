package common

import (
	"time"
	"bisale/foundation/thrift/pool"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"bisale/bisale-console-api/thrift/balanceAccount"
)
func openBisaleBalanceAccountServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "chainDepositWithdraw")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale balance account service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := balanceAccount.NewTChainDepositWithdrawServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleBalanceAccountServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale balance account service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale balance account client success")
	return nil
}

func GetBisaleBalanceAccountServiceClient() (s *balanceAccount.TChainDepositWithdrawServiceClient, c *thriftPool.IdleClient) {

	client, err := BisaleBalanceAccountServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale balance account service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale balance account client has closed"))
	}

	Log.Info("Get bisale business client from pool success")

	return client.Client.(*balanceAccount.TChainDepositWithdrawServiceClient), client
}