// Code generated by protoc-gen-go.
// source: geoip2/geoip2.proto
// DO NOT EDIT!

/*
Package geoip2 is a generated protocol buffer package.

It is generated from these files:
	geoip2/geoip2.proto

It has these top-level messages:
	IpRequest
	GeoInfo
	City
	Continent
	Country
	Location
	Postal
	RepresentedCountry
	Traits
*/
package geoip2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The request with ip address
type IpRequest struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *IpRequest) Reset()         { *m = IpRequest{} }
func (m *IpRequest) String() string { return proto.CompactTextString(m) }
func (*IpRequest) ProtoMessage()    {}

// The response message containing the Geo info about current IP address
type GeoInfo struct {
	City               *City               `protobuf:"bytes,1,opt,name=city" json:"city,omitempty"`
	Continent          *Continent          `protobuf:"bytes,2,opt,name=continent" json:"continent,omitempty"`
	Country            *Country            `protobuf:"bytes,3,opt,name=country" json:"country,omitempty"`
	Location           *Location           `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
	Postal             *Postal             `protobuf:"bytes,5,opt,name=postal" json:"postal,omitempty"`
	RegisteredCountry  *Country            `protobuf:"bytes,6,opt,name=registeredCountry" json:"registeredCountry,omitempty"`
	RepresentedCountry *RepresentedCountry `protobuf:"bytes,7,opt,name=representedCountry" json:"representedCountry,omitempty"`
	Subdivisions       []*Country          `protobuf:"bytes,8,rep,name=subdivisions" json:"subdivisions,omitempty"`
	Traits             *Traits             `protobuf:"bytes,9,opt,name=traits" json:"traits,omitempty"`
}

func (m *GeoInfo) Reset()         { *m = GeoInfo{} }
func (m *GeoInfo) String() string { return proto.CompactTextString(m) }
func (*GeoInfo) ProtoMessage()    {}

func (m *GeoInfo) GetCity() *City {
	if m != nil {
		return m.City
	}
	return nil
}

func (m *GeoInfo) GetContinent() *Continent {
	if m != nil {
		return m.Continent
	}
	return nil
}

func (m *GeoInfo) GetCountry() *Country {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *GeoInfo) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *GeoInfo) GetPostal() *Postal {
	if m != nil {
		return m.Postal
	}
	return nil
}

func (m *GeoInfo) GetRegisteredCountry() *Country {
	if m != nil {
		return m.RegisteredCountry
	}
	return nil
}

func (m *GeoInfo) GetRepresentedCountry() *RepresentedCountry {
	if m != nil {
		return m.RepresentedCountry
	}
	return nil
}

func (m *GeoInfo) GetSubdivisions() []*Country {
	if m != nil {
		return m.Subdivisions
	}
	return nil
}

func (m *GeoInfo) GetTraits() *Traits {
	if m != nil {
		return m.Traits
	}
	return nil
}

// type of City
type City struct {
	GeoNameID uint32            `protobuf:"varint,1,opt,name=geoNameID" json:"geoNameID,omitempty"`
	Names     map[string]string `protobuf:"bytes,2,rep,name=names" json:"names,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *City) Reset()         { *m = City{} }
func (m *City) String() string { return proto.CompactTextString(m) }
func (*City) ProtoMessage()    {}

func (m *City) GetNames() map[string]string {
	if m != nil {
		return m.Names
	}
	return nil
}

type Continent struct {
	Code      string            `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	GeoNameID uint32            `protobuf:"varint,2,opt,name=geoNameID" json:"geoNameID,omitempty"`
	Names     map[string]string `protobuf:"bytes,3,rep,name=names" json:"names,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Continent) Reset()         { *m = Continent{} }
func (m *Continent) String() string { return proto.CompactTextString(m) }
func (*Continent) ProtoMessage()    {}

func (m *Continent) GetNames() map[string]string {
	if m != nil {
		return m.Names
	}
	return nil
}

type Country struct {
	IsoCode   string            `protobuf:"bytes,1,opt,name=isoCode" json:"isoCode,omitempty"`
	GeoNameID uint32            `protobuf:"varint,2,opt,name=geoNameID" json:"geoNameID,omitempty"`
	Names     map[string]string `protobuf:"bytes,3,rep,name=names" json:"names,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Country) Reset()         { *m = Country{} }
func (m *Country) String() string { return proto.CompactTextString(m) }
func (*Country) ProtoMessage()    {}

func (m *Country) GetNames() map[string]string {
	if m != nil {
		return m.Names
	}
	return nil
}

type Location struct {
	Latitude  int64  `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude int64  `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
	MetroCode uint32 `protobuf:"varint,3,opt,name=metroCode" json:"metroCode,omitempty"`
	TimeZone  string `protobuf:"bytes,4,opt,name=timeZone" json:"timeZone,omitempty"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}

type Postal struct {
	Code string `protobuf:"bytes,1,opt,name=Code" json:"Code,omitempty"`
}

func (m *Postal) Reset()         { *m = Postal{} }
func (m *Postal) String() string { return proto.CompactTextString(m) }
func (*Postal) ProtoMessage()    {}

type RepresentedCountry struct {
	IsoCode   string            `protobuf:"bytes,1,opt,name=isoCode" json:"isoCode,omitempty"`
	GeoNameID uint32            `protobuf:"varint,2,opt,name=geoNameID" json:"geoNameID,omitempty"`
	Names     map[string]string `protobuf:"bytes,3,rep,name=names" json:"names,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Type      string            `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
}

func (m *RepresentedCountry) Reset()         { *m = RepresentedCountry{} }
func (m *RepresentedCountry) String() string { return proto.CompactTextString(m) }
func (*RepresentedCountry) ProtoMessage()    {}

func (m *RepresentedCountry) GetNames() map[string]string {
	if m != nil {
		return m.Names
	}
	return nil
}

type Traits struct {
	IsAnonymousProxy    bool `protobuf:"varint,1,opt,name=isAnonymousProxy" json:"isAnonymousProxy,omitempty"`
	IsSatelliteProvider bool `protobuf:"varint,2,opt,name=isSatelliteProvider" json:"isSatelliteProvider,omitempty"`
}

func (m *Traits) Reset()         { *m = Traits{} }
func (m *Traits) String() string { return proto.CompactTextString(m) }
func (*Traits) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for GeoIP service

type GeoIPClient interface {
	Lookup(ctx context.Context, in *IpRequest, opts ...grpc.CallOption) (*GeoInfo, error)
}

type geoIPClient struct {
	cc *grpc.ClientConn
}

func NewGeoIPClient(cc *grpc.ClientConn) GeoIPClient {
	return &geoIPClient{cc}
}

func (c *geoIPClient) Lookup(ctx context.Context, in *IpRequest, opts ...grpc.CallOption) (*GeoInfo, error) {
	out := new(GeoInfo)
	err := grpc.Invoke(ctx, "/geoip2.GeoIP/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GeoIP service

type GeoIPServer interface {
	Lookup(context.Context, *IpRequest) (*GeoInfo, error)
}

func RegisterGeoIPServer(s *grpc.Server, srv GeoIPServer) {
	s.RegisterService(&_GeoIP_serviceDesc, srv)
}

func _GeoIP_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(IpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GeoIPServer).Lookup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _GeoIP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "geoip2.GeoIP",
	HandlerType: (*GeoIPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lookup",
			Handler:    _GeoIP_Lookup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
