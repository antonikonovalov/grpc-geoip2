FROM google/golang

ADD . /gopath/src/github.com/antonikonovalov/grpc-geoip2
WORKDIR /gopath/src/github.com/antonikonovalov/grpc-geoip2
RUN go get github.com/antonikonovalov/grpc-geoip2
RUN make
RUN export PATH=$PATH:/gopath/bin

EXPOSE 50052

ENTRYPOINT ["/gopath/bin/geoip2-server"]