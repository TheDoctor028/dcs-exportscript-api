package api

import udpConnection "github.com/thedoctor028/dcsexportscriptapi/udp-connection"

type API struct {
	udpSender *udpConnection.UDPSender
}
