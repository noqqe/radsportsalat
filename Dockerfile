FROM golang:1.17

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get -v ./...
RUN go build .

# Run radsportsalat
EXPOSE 8080
CMD [ "./radsportsalat" ]
