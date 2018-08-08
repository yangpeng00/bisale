package common

import (
	"time"
	"bisale/foundation/thrift/pool"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"bisale/bisale-console-api/thrift/content"
)
func openBisaleAppVersionServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	protocol := thrift.NewTBinaryProtocolTransport(transport)

	mp := thrift.NewTMultiplexedProtocol(protocol, "appVersion")

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open bisale appVersion service connection error: %s", err.Error()))
		return nil, err
	}

	standardClient := thrift.NewTStandardClient(mp, mp)

	client := content.NewTAppVersionServiceClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeBisaleAppVersionServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close bisale appVersion service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close bisale appVersion client success")
	return nil
}

func GetBisaleAppVersionServiceClient() (s *content.TAppVersionServiceClient, c *thriftPool.IdleClient) {

	client, err := BisaleAppVersionServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get bisale appVersion service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Bisale appVersion client has closed"))
	}

	Log.Info("Get bisale business client from pool success")

	return client.Client.(*content.TAppVersionServiceClient), client
}