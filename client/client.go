package client

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/antonikonovalov/grpc-geoip2/geoip2"
)

const (
	DefaultAddress = "0.0.0.0:50052"
)

type Client struct {
	grpc pb.GeoIPClient
	cc   *grpc.ClientConn
}

//create new connection and configure client
func New(address string) (*Client, error) {
	if len(address) == 0 {
		address = DefaultAddress
	}
	client := new(Client)
	// Set up a connection to the server.
	conn, err := grpc.Dial(address)
	if err != nil {
		return nil, err
	}
	client.cc = conn
	client.grpc = pb.NewGeoIPClient(conn)
	return client, nil
}

func (c *Client) Lookup(ip string) (*pb.GeoInfo, error) {
	return c.grpc.Lookup(context.Background(), &pb.IpRequest{Ip: ip})
}

//close connection with server
func (c *Client) Close() error {
	return c.cc.Close()
}
