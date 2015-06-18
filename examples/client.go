package main

import (
	"flag"
	"fmt"

	"github.com/antonikonovalov/grpc-geoip2/client"
)

var (
	ip = flag.String("ip", "101.88.10.12", "target search")
)

func main() {
	flag.Parse()
	geoip2, err := client.New("")
	if err != nil {
		fmt.Printf("Can't create client: %s", err)
	}
	defer geoip2.Close()
	r, err := geoip2.Lookup(*ip)
	if err != nil {
		fmt.Printf("Can't lookup: %s\n", err)
	} else {
		if r.City.GetNames() != nil {
			fmt.Printf("\n====================\nEnglish city name: %v\n", r.City.Names["en"])
		}
		if r.GetSubdivisions() != nil {
			fmt.Printf("English subdivision name: %v\n", r.Subdivisions[0].Names["en"])
		}

		if r.GetCountry() != nil {
			fmt.Printf("Russian country name: %v\n", r.Country.Names["ru"])
			fmt.Printf("ISO country code: %v\n", r.Country.IsoCode)
		}
		if r.GetLocation() != nil {
			fmt.Printf("Time zone: %v\n", r.Location.TimeZone)
			fmt.Printf("Coordinates: %v, %v\n", r.Location.Latitude, r.Location.Longitude)
		}
	}
}
