package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestData = `192.168.100.241,64453,178.248.237.68,443(https),3
192.168.100.241,64455,178.248.237.68,443(https),3
178.248.237.68,443(https),192.168.100.241,64455,3
192.168.100.241,64471,216.58.211.17,443(https),3
`

func TestProcess(t *testing.T) {
	sniff, err := NewSniffer("packet_test.pcap")
	if err != nil {
		t.Error(err)
	}

	var messages string
	for mes := range sniff.Run() {
		messages += mes + "\n"
	}

	assert.Equal(t, strings.Compare(TestData, messages), 0)
}
