FROM golang:1.17

RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main .

CMD ["./main"]