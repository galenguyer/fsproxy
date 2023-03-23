FROM docker.io/golang:1.20-alpine AS builder

WORKDIR /src/
COPY go.mod go.sum main.go /src/
RUN go build

FROM alpine
COPY --from=builder /src/lolproxy /lolproxy
EXPOSE 6969
ENTRYPOINT [ "/lolproxy", "--host", "0.0.0.0:6969" ]
