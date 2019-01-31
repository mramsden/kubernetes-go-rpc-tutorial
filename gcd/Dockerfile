FROM golang:1.11

WORKDIR /go/src/github.com/mramsden/kubernetes-go-grpc-tutorial/gcd
COPY gcd .
COPY pb ../pb

RUN go get -v ./...
RUN go install -v ./...

EXPOSE 3000

CMD [ "gcd" ]
