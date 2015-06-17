package main

import (
	"log"
	"net"

	"flag"
	pb "github.com/antonikonovalov/grpc-geoip2/geoip2"
	"github.com/oschwald/geoip2-golang"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	db       *geoip2.Reader
	port     = flag.String("port", ":50052", "port for listen service (default - :50052)")
	pathToDb = flag.String("db", "GeoIP2-City.mmdb", "path to db for lookup data by ip (default - GeoIP2-City.mmdb)")
)

// server is used to implement geoip2.GeoIPServer
type server struct{}

// Lookup implements geoip2.GeoIPServer
func (s *server) Lookup(ctx context.Context, in *pb.IpRequest) (*pb.GeoInfo, error) {
	ip := net.ParseIP(in.Ip)
	record, err := db.City(ip)
	if err != nil {
		log.Print("Error: ", err)
		return nil, err
	}

	return cityToGeoInfo(record), nil
}

func cityToGeoInfo(city *geoip2.City) *pb.GeoInfo {
	info := &pb.GeoInfo{}
	//city
	info.City.GeoNameID = uint32(city.City.GeoNameID)
	info.City.Names = city.City.Names
	//Continent
	info.Continent.GeoNameID = uint32(city.Continent.GeoNameID)
	info.Continent.Code = city.Continent.Code
	info.Continent.Names = city.Continent.Names
	//Country
	info.Country.IsoCode = city.Country.IsoCode
	info.Country.Names = city.Country.Names
	info.Country.GeoNameID = uint32(city.Country.GeoNameID)
	//postal
	info.Postal.Code = city.Postal.Code
	//location
	info.Location.MetroCode = uint32(city.Location.MetroCode)
	info.Location.TimeZone = city.Location.TimeZone
	info.Location.Latitude = int64(city.Location.Latitude * 1000000.0)
	info.Location.Longitude = int64(city.Location.Longitude * 1000000.0)
	//RegisteredCountry
	info.RegisteredCountry.IsoCode = city.RegisteredCountry.IsoCode
	info.RegisteredCountry.Names = city.RegisteredCountry.Names
	info.RegisteredCountry.GeoNameID = uint32(city.RegisteredCountry.GeoNameID)
	//RegisteredCountry
	info.RepresentedCountry.IsoCode = city.RepresentedCountry.IsoCode
	info.RepresentedCountry.Names = city.RepresentedCountry.Names
	info.RepresentedCountry.GeoNameID = uint32(city.RepresentedCountry.GeoNameID)
	info.RepresentedCountry.Type = city.RepresentedCountry.Type
	//Subdivisions
	for _, subdiv := range city.Subdivisions {
		info.Subdivisions = append(info.Subdivisions, &pb.Country{
			Names:     subdiv.Names,
			GeoNameID: uint32(subdiv.GeoNameID),
			IsoCode:   subdiv.IsoCode,
		})
	}
	//Traits
	info.Traits.IsAnonymousProxy = city.Traits.IsAnonymousProxy
	info.Traits.IsSatelliteProvider = city.Traits.IsSatelliteProvider

	return info
}

func main() {
	flag.Parse()
	var err error
	db, err = geoip2.Open(*pathToDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGeoIPServer(s, &server{})
	log.Print("serve 0.0.0.0"+*port+" db from "+*pathToDb)
	s.Serve(lis)
}
