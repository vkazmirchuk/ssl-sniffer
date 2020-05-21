# Build
FROM golang:latest as builder

RUN echo "deb http://ftp.de.debian.org/debian bullseye main" >> /etc/apt/sources.list &&\
    apt update &&\
    apt install libpcap-dev -y

ENV GOOS=linux GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o /build/main .

EXPOSE 8080

ENTRYPOINT ["/build/main", "packet_test.pcap"]
