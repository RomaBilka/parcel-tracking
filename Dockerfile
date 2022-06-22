FROM golang

WORKDIR /go/src/app
COPY . .

CMD ["go", "run", "cmd/parcel-tracking/main.go"]