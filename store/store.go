package store

import "github.com/antonikonovalov/grpc-geoip2/geoip2"

type Store interface {
	//find geo data by ip
	Lookup(ip *geoip2.IpRequest) (*geoip2.GeoInfo, error)
	//close all connection others stores
	Close() error
}
