# Build binary
FROM golang:1.16-alpine as build

ARG VERSION="-"
ARG COMMIT="-"
ARG DATE="-"

WORKDIR /go/src/github.com/prune998/gojsontoenv
ADD . /go/src/github.com/prune998/gojsontoenv

RUN go get -d -v ./...
RUN go build -ldflags="-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}" -o /go/bin/gojsontoenv

# Create final image
FROM gcr.io/distroless/base

USER nonroot
COPY --from=build /go/bin/gojsontoenv /

ENTRYPOINT ["/gojsontoenv"]
