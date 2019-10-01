# stage 0
FROM golang:1.13.1-buster as builder
WORKDIR /build
COPY ./ /build
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -ldflags '-s -linkmode external -extldflags "-static"' 

# stage 1
FROM scratch
WORKDIR /

COPY --from=builder /build/sentry_scratch_golang .

ENTRYPOINT ["/sentry_scratch_golang"]