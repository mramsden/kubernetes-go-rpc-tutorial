FROM golang:1.11

WORKDIR /go/src/github.com/mramsden/kubernetes-go-grpc-tutorial/api

COPY api .
COPY pb ../pb

RUN go get -v ./...
RUN go install -v ./...

EXPOSE 3000

CMD [ "api" ]
