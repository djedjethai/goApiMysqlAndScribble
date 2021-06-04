FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

WORKDIR /go/src/app/cmd/beer-server
RUN go build main.go


CMD ["./main"]
