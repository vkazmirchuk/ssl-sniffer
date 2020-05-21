package main

import (
	"fmt"
	"io"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Sniffer struct {
	handle *pcap.Handle
}

// Create new sniffer
func NewSniffer(src string) (sniff *Sniffer, err error) {
	sniff = &Sniffer{}

	// Check src file exists
	if _, err := os.Stat(src); err == nil {

		// Open file
		file, err := os.Open(src)
		if err != nil {
			return nil, err
		}
		sniff.handle, err = pcap.OpenOfflineFile(file)
		if err != nil {
			return nil, err
		}

	} else {

		// Open device
		sniff.handle, err = pcap.OpenLive(src, 1600, false, pcap.BlockForever)
		if err != nil {
			return nil, err
		}

	}

	// Set filter
	if err := sniff.SetFilter(); err != nil {
		return nil, err
	}

	return
}

// Set BPF filter
func (s *Sniffer) SetFilter() error {
	// filter := "tcp port 443 and (tcp[((tcp[12] & 0xf0) >> 2)] = 0x15)"
	filter := "tcp port 443 and (tcp[((tcp[12] & 0xf0) >> 2)] = 0x16)"
	return s.handle.SetBPFFilter(filter)
}

func (s *Sniffer) Run() (c chan string) {
	c = make(chan string)
	go s.Process(c)
	return
}

func (s *Sniffer) Process(c chan string) {
	// Set parser
	var eth layers.Ethernet
	var ip4 layers.IPv4
	var tcp layers.TCP
	var tls layers.TLS

	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ip4, &tcp, &tls)

	decodedLayers := make([]gopacket.LayerType, 0, 5)

	for {
		// Read packet
		data, _, err := s.handle.ReadPacketData()
		if err != nil {
			if err == io.EOF {
				close(c)
				break
			}
			fmt.Println("[ERROR read packet]", err)
			continue
		}

		var IP_SRC, TCP_SRC, IP_DST, TCP_DST string
		var TCP_OPTIONS_COUNT int

		// Decode layers
		if err := parser.DecodeLayers(data, &decodedLayers); err != nil {
			continue
		}
		for _, typ := range decodedLayers {
			switch typ {
			case layers.LayerTypeIPv4:
				IP_SRC = ip4.SrcIP.String()
				IP_DST = ip4.DstIP.String()
			case layers.LayerTypeTCP:
				TCP_SRC = tcp.SrcPort.String()
				TCP_DST = tcp.DstPort.String()
				TCP_OPTIONS_COUNT = len(tcp.Options)
			}
		}

		message := fmt.Sprintf("%s,%s,%s,%s,%d", IP_SRC, TCP_SRC, IP_DST, TCP_DST, TCP_OPTIONS_COUNT)
		c <- message
	}
}
