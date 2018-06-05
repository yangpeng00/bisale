package common

import (
	"net"
	"fmt"
	"time"
	"bisale/foundation/thrift/pool"
	"bisale/bisale-console-api/thrift/captcha"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func openCaptchaServiceClient(host, port string, ConnTimeout time.Duration) (*thriftPool.IdleClient, error) {

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, _ := thrift.NewTSocket(net.JoinHostPort(host, port))

	transport := thrift.NewTFramedTransport(socket)

	if err := transport.Open(); err != nil {
		Log.Error(fmt.Printf("Open captcha service connection error: %s", err.Error()))
		return nil, err
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	standardClient := thrift.NewTStandardClient(iprot, oprot)

	client := captcha.NewCaptchaClient(standardClient)

	return &thriftPool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func closeCaptchaServiceClient(c *thriftPool.IdleClient) error {
	err := c.Socket.Close()
	if err != nil {
		Log.Error(fmt.Printf("Close captcha service connection error: %s", err.Error()))
		return err
	}
	Log.Info("Close captcha client success")
	return nil
}

func GetCaptchaServiceClient() (s *captcha.CaptchaClient, c *thriftPool.IdleClient) {

	client, err := CaptchaServicePool.Get()

	if err != nil {
		Log.Error(fmt.Printf("Get captcha service client error: %s", err.Error()))
		return
	}

	if !client.Socket.IsOpen() {
		Log.Error(fmt.Printf("Captcha client has closed"))
	}

	//err = CaptchaServicePool.Put(client)
	//
	//if err != nil {
	//	Log.Error(fmt.Printf("Put captcha client to pool error: %s", err.Error()))
	//	return
	//}

	Log.Info("Get captcha client from pool success")

	return client.Client.(*captcha.CaptchaClient), client
}
