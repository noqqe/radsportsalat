FROM golang:1.21-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /go/src/app
COPY templates /go/src/app/templates
COPY assets /go/src/app/assets
COPY ical /go/src/app/ical
COPY go.mod /go/src/app/go.mod
COPY go.sum /go/src/app/go.sum
COPY .git /go/src/app/.git
COPY radsportsalat.go /go/src/app/radsportsalat.go
COPY data.go /go/src/app/data.go

# build
#RUN go get -v ./...
RUN go build -ldflags "-X main.Version=`git describe --tags`"  -v .

# copy
FROM scratch
WORKDIR /go/src/app
COPY --from=builder /go/src/app/radsportsalat /go/src/app/radsportsalat
COPY templates /go/src/app/templates
COPY ical /go/src/app/ical
COPY assets /go/src/app/assets

# run
EXPOSE 8080
CMD [ "/go/src/app/radsportsalat" ]
