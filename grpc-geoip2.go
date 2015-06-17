package main

import (
	"log"
	"net"

	"flag"
	pb "github.com/antonikonovalov/grpc-geoip2/geoip2"
	"github.com/oschwald/geoip2-golang"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/proto"
	"github.com/boltdb/bolt"
)

var (
	db       *geoip2.Reader
	cache    *bolt.DB
	port     = flag.String("port", ":50051", "port for listen service (default - :50051)")
	pathToDb = flag.String("db", "GeoIP2-City.mmdb", "path to db for lookup data by ip (default - GeoIP2-City.mmdb)")
)

// server is used to implement geoip2.GeoIPServer
type server struct{}

// Lookup implements geoip2.GeoIPServer
func (s *server) Lookup(ctx context.Context, in *pb.IpRequest) (*pb.GeoInfo, error) {
	ip := net.ParseIP(in.Ip)
	info := &pb.GeoInfo{}
	err := cache.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ips"))
		v := b.Get([]byte(ip))
		if len(v) == 0 {
			return nil
		}
		//log.Printf("The answer is: %s\n", v)
		return proto.Unmarshal(v,info)
	})
	if err != nil {
		return nil,err
	}
	if info.GetCountry() != nil {
		log.Print("return from cache")
		return info,nil
	}

	record, err := db.City(ip)
	if err != nil {
		log.Print("Error: ", err)
		return nil, err
	}
	info = cityToGeoInfo(record)
	err = cache.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ips"))
		bytes,err := proto.Marshal(info)
		if err != nil {
			return err
		}
		err = b.Put([]byte(ip), bytes)
		return err
	})

	if err != nil {
		return nil,err
	}

	log.Print("return from db")
	return info, nil
}

func cityToGeoInfo(city *geoip2.City) *pb.GeoInfo {
	info := &pb.GeoInfo{}
	//city
	info.City = &pb.City{
		GeoNameID: uint32(city.City.GeoNameID),
		Names:     city.City.Names,
	}
	//Continent
	info.Continent = &pb.Continent{
		GeoNameID: uint32(city.Continent.GeoNameID),
		Code:      city.Continent.Code,
		Names:     city.Continent.Names,
	}
	//Country
	info.Country = &pb.Country{
		IsoCode:   city.Country.IsoCode,
		Names:     city.Country.Names,
		GeoNameID: uint32(city.Country.GeoNameID),
	}
	//postal
	info.Postal = &pb.Postal{Code: city.Postal.Code}
	//location
	info.Location = &pb.Location{
		MetroCode: uint32(city.Location.MetroCode),
		TimeZone:  city.Location.TimeZone,
		Latitude:  int64(city.Location.Latitude * 1000000.0),
		Longitude: int64(city.Location.Longitude * 1000000.0),
	}
	//RegisteredCountry
	info.RegisteredCountry = &pb.Country{
		IsoCode:   city.RegisteredCountry.IsoCode,
		Names:     city.RegisteredCountry.Names,
		GeoNameID: uint32(city.RegisteredCountry.GeoNameID),
	}
	//RegisteredCountry
	info.RepresentedCountry = &pb.RepresentedCountry{
		IsoCode:   city.RepresentedCountry.IsoCode,
		Names:     city.RepresentedCountry.Names,
		GeoNameID: uint32(city.RepresentedCountry.GeoNameID),
		Type:      city.RepresentedCountry.Type,
	}
	//Subdivisions
	for _, subdiv := range city.Subdivisions {
		info.Subdivisions = append(info.Subdivisions, &pb.Country{
			Names:     subdiv.Names,
			GeoNameID: uint32(subdiv.GeoNameID),
			IsoCode:   subdiv.IsoCode,
		})
	}
	//Traits
	info.Traits = &pb.Traits{
		IsAnonymousProxy:    city.Traits.IsAnonymousProxy,
		IsSatelliteProvider: city.Traits.IsSatelliteProvider,
	}

	return info
}

func main() {
	flag.Parse()
	var err error
	//max db
	db, err = geoip2.Open(*pathToDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//bolt db for cache
	cache, err = bolt.Open("cache.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.Close()

	cache.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("ips"))
		if err != nil {
			log.Fatal("create bucket:", err)
			return err
		}
		return nil
	})

	//tray listen
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGeoIPServer(s, &server{})
	log.Print("serve 0.0.0.0" + *port + " db from " + *pathToDb)
	s.Serve(lis)
}
