# Build binary
FROM golang:1.16-alpine as build

WORKDIR /go/src/github.com/prune998/gojsontoenv
ADD . /go/src/github.com/prune998/gojsontoenv

RUN go get -d -v ./...
RUN go build -o /go/bin/gojsontoenv

# Create final image
FROM gcr.io/distroless/base

USER nonroot
COPY --from=build /go/bin/gojsontoenv /

ENTRYPOINT ["/gojsontoenv"]
