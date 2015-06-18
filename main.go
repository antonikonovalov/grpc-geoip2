package main

import (
	"flag"
	"github.com/antonikonovalov/grpc-geoip2/store/database"
	"github.com/antonikonovalov/grpc-geoip2/server"
)

var (
	port     = flag.String("port", ":50052", "port for listen service (default - :50052)")
	mmdbPath = flag.String("mmdb", "", "path to db for lookup data by ip (default - "+database.DefaultPathMmdb+")")
	cachePath = flag.String("cache", "", "path to db for cache data after find (default - "+database.DefaultPathCache+")")
)

func main() {
	flag.Parse()
	//if err - fatal in newStore
	store := database.NewStore(*mmdbPath,*cachePath)
	server.New(*port,store).Serve()
}
