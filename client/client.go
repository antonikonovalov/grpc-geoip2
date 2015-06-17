package main

import (
	"log"

	pb "github.com/antonikonovalov/grpc-geoip2/geoip2"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"flag"
	"fmt"
)

var (
	address     = "0.0.0.0:50051"
	ip = flag.String("ip","81.2.69.142","set ip for lookup")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGeoIPClient(conn)

	// Contact the server and print out its response.


	r, err := c.Lookup(context.Background(), &pb.IpRequest{Ip: *ip})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}
	//log.Printf("Lookuped: %s", r.String())
	fmt.Printf("\n\nEnglish city name: %v\n", r.City.Names["en"])
	fmt.Printf("English subdivision name: %v\n", r.Subdivisions[0].Names["en"])
	fmt.Printf("Russian country name: %v\n", r.Country.Names["ru"])
	fmt.Printf("ISO country code: %v\n", r.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", r.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", r.Location.Latitude, r.Location.Longitude)
}