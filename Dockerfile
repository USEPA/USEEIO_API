FROM golang
ADD ./src /go/src/epa/useeio-api
RUN go install epa/useeio-api
EXPOSE 80
ENTRYPOINT /go/bin/useeio-api -data "/go/src/epa/useeio-api/data" -port 80