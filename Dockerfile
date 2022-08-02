# Builder stage
FROM golang:1.18-alpine as builder

WORKDIR /go/src/app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/rest-server ./cmd/rest-server/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /go/bin/rest-server .

USER 65532:65532
EXPOSE 8080

CMD ["/rest-server"]
