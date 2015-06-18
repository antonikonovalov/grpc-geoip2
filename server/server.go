package server

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/antonikonovalov/grpc-geoip2/geoip2"
	"github.com/antonikonovalov/grpc-geoip2/store"

)

// server is used to implement geoip2.GeoIPServer
type Server struct {
	port  string
	store store.Store
}

func New(port string, store store.Store) *Server {
	return &Server{port: port, store: store}
}

func (s *Server) Serve() {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	//close store after finish serve
	defer s.store.Close()
	gs := grpc.NewServer()
	pb.RegisterGeoIPServer(gs, s)
	println("serve 0.0.0.0" + s.port)
	gs.Serve(lis)
}

// Lookup implements geoip2.GeoIPServer
func (s *Server) Lookup(ctx context.Context, ip *pb.IpRequest) (*pb.GeoInfo, error) {
	//find in store
	return s.store.Lookup(ip)
}