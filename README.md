# grpc-geoip2

GeoIP2 like web service with grpc

# install

## From Source 

* install Golang 
*  `go get github.com/antonikonovalov/grpc-geoip2`
* download GeoLite2-City.mmdb -> /db/geoip2/GeoLite2-City.mmdb
* run and enjoy! 

```
    >grpc-geoip2 -mmdb=/path/to/my.mmdb
    serve 0.0.0.0:50052
    >go run example/client.go -ip=123.23.34.56
```
