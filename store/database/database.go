package database

import (
	"fmt"
	"log"
	"net"

	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
	"github.com/oschwald/geoip2-golang"

	pb "github.com/antonikonovalov/grpc-geoip2/geoip2"
	"github.com/antonikonovalov/grpc-geoip2/store"
	"errors"
)

const (
	DefaultPathMmdb  = `/db/geoip2/GeoLite2-City.mmdb`
	DefaultPathCache = `/db/geoip2/cache.db`
)

func NewStore(mmdbPath, cachePath string) store.Store {
	var err error
	store := new(Store)

	if len(mmdbPath) == 0 {
		mmdbPath = DefaultPathMmdb
	}
	if len(cachePath) == 0 {
		cachePath = DefaultPathCache
	}
	store.mmdb, err = geoip2.Open(mmdbPath)
	if err != nil {
		log.Fatal(err)
	}
	store.cache, err = bolt.Open("cache.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = store.cache.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("ips"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal("create bucket:", err)
	}

	return store
}

type Store struct {
	mmdb  *geoip2.Reader
	cache *bolt.DB
}

func (s *Store) Close() error {
	//without error
	s.mmdb.Close()
	return s.cache.Close()
}

func (s *Store) Lookup(in *pb.IpRequest) (*pb.GeoInfo,error) {
	ip := net.ParseIP(in.Ip)
	if ip == nil {
		return nil,errors.New("invalid ip")
	}
	info := &pb.GeoInfo{}
	err := s.cache.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ips"))
		v := b.Get([]byte(ip))
		if len(v) == 0 {
			return nil
		}
		//log.Printf("The answer is: %s\n", v)
		return proto.Unmarshal(v, info)
	})
	if err != nil {
		return nil, err
	}
	if info.GetCountry() != nil {
		return info, nil
	}

	record, err := s.mmdb.City(ip)
	if err != nil {
		fmt.Errorf("Error: %s", err)
		return nil, err
	}
	info = cityToGeoInfo(record)
	err = s.cache.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ips"))
		bytes, err := proto.Marshal(info)
		if err != nil {
			return err
		}
		err = b.Put([]byte(ip), bytes)
		return err
	})
	if err != nil {
		return nil, err
	}
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
