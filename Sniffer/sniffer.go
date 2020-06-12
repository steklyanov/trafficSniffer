package Sniffer

import (
	"fmt"
	"github.com/bradleyfalzon/tlsx"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"time"
)

func ReadPacket(packet gopacket.Packet) {
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, ok := tcpLayer.(*layers.TCP)
		if !ok {
			log.Println("Could not decode TCP layer")
			return
		}
		if tcp.SYN {
			// Connection setup
		} else if tcp.FIN {
			// Connection teardown
		} else if tcp.ACK && len(tcp.LayerPayload()) == 0 {
			// Acknowledgement packet
		} else if tcp.RST {
			// Unexpected packet
		} else {
			// data packet
			readData(packet)
		}
	}
}

func FindDevices()  {
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Print device information
	fmt.Println("Devices found:")
	for _, device := range devices {
		listen(device.Name)
	}
}

var (
	snapshotLen int32  = 1024
	promiscuous  bool   = false
	timeout      time.Duration = -1 * time.Second
)

func listen(device string){
		handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
		if err != nil {
			log.Println(err)
			return
		}
		defer handle.Close()
		err = handle.SetBPFFilter("(port 443)")
		if err != nil {
			log.Fatal(err)
		}

		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			go readPacket(packet)
		}
}

func readPacket(packet gopacket.Packet) {
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, ok := tcpLayer.(*layers.TCP)
		if !ok {
			log.Println("Could not decode TCP layer")
			return
		}
		if tcp.SYN {
			// Connection setup
		} else if tcp.FIN {
			// Connection teardown
		} else if tcp.ACK && len(tcp.LayerPayload()) == 0 {
			// Acknowledgement packet
		} else if tcp.RST {
			// Unexpected packet
		} else {
			// data packet
			readData(packet)
		}
	}
}

func readData(packet gopacket.Packet) {
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		t, _ := tcpLayer.(*layers.TCP)
		var hello = tlsx.ClientHello{}
		var ipDst  net.IP
		var ipSrc net.IP

		err := hello.Unmarshall(t.LayerPayload())

		switch err {
		case nil:
		case tlsx.ErrHandshakeWrongType:
			return
		default:
			log.Println("Error reading Client Hello:", err)
			log.Println("Raw Client Hello:", t.LayerPayload())
			return
		}
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			ipDst = ip.DstIP
			ipSrc = ip.SrcIP
		} else {
			ipLayer = packet.Layer(layers.LayerTypeIPv6)
			ip, _ := ipLayer.(*layers.IPv4)
			ipDst = ip.DstIP
			ipSrc = ip.SrcIP
		}
		fmt.Println(ipSrc, t.SrcPort, ipDst, t.DstPort)
	} else { return }
}