package common

import (
	"time"
	"bisale/foundation/thrift/pool"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"bisale/bisale-console-api/thrift/finance"
)
func openBisaleAccountStatementServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "accountStatement")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale account statement service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := finance.NewTAccountStatementServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleAccountStatementServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale account statement service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale account statement client success")
	return nil
}

func GetBisaleAccountStatementServiceClient() (s *finance.TAccountStatementServiceClient, c *thriftPool.IdleClient) {

	client, err := BisaleAccountStatementPool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale accountStatement service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale accountStatement client has closed"))
	}

	Log.Info("Get bisale business client from pool success")

	return client.Client.(*finance.TAccountStatementServiceClient), client
}