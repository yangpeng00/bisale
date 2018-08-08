package common

import (
	"time"
	"bisale/foundation/thrift/pool"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"bisale/bisale-console-api/thrift/engine"
)
func openBisaleTradeDetailServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "trade")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale tradeDetail service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := engine.NewTTradeServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleTradeDetailServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale tradeDetail service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale tradeDetail client success")
	return nil
}

func GetBisaleTradeDetailServiceClient() (s *engine.TTradeServiceClient, c *thriftPool.IdleClient) {

	client, err := BisaleTradeDetailServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale tradeDetail service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale tradeDetail client has closed"))
	}

	Log.Info("Get bisale business client from pool success")

	return client.Client.(*engine.TTradeServiceClient), client
}