FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on

COPY ./ /src

WORKDIR /src
RUN GOGC=off go build -v -o /power-datacenter-exporter .

FROM gcr.io/distroless/static-debian11

COPY --from=builder /power-datacenter-exporter /power-datacenter-exporter

ENTRYPOINT  [ "/power-datacenter-exporter" ]
EXPOSE 8080
