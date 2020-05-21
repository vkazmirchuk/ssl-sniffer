# Sniffer SSL handshake packets
- Run on a 64-bit Linux distribution (Centos, Ubuntu, Debian).
- Sniff tcp/ip packets.
- Detect among the sniffed packets detect SSL (https) handshake packets.
- Print to stdout each detection in the following format: IP_SRC,TCP_SRC,IP_DST,TCP_DST,COUNT(TCP_OPTIONS).

## How to launch

### Run tests
`go test`

### Build
`GOOS=linux GOARCH=amd64 go build -o sniffer .`

### Run
`./sniffer eth1` (`eth1` is an interface) for online sniff

or 

`./sniffer dump.pcap` (`dump.pcap` is an dump file) for offline file sniff

Check: http://127.0.0.1:8080/

## How to launch in Docker
Only file analysis will work in Docker

### Pull image
`docker pull webdizi/ssl-sniffer`

### Run
`docker run --rm -i -t -v $(pwd)/myDump.pcap:/build/packet_test.pcap webdizi/ssl-sniffer:latest` (`myDump.pcap` is your dump file)

