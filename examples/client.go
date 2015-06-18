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
	r, err := geoip2.Lookup(*ip)
	if err != nil {
		fmt.Printf("Can't lookup: %s\n", err)
	} else {
		fmt.Printf("\nEnglish city name: %v\n", r.City.Names["en"])
		fmt.Printf("English subdivision name: %v\n", r.Subdivisions[0].Names["en"])
		fmt.Printf("Russian country name: %v\n", r.Country.Names["ru"])
		fmt.Printf("ISO country code: %v\n", r.Country.IsoCode)
		fmt.Printf("Time zone: %v\n", r.Location.TimeZone)
		fmt.Printf("Coordinates: %v, %v\n", r.Location.Latitude, r.Location.Longitude)
	}
}
