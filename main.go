/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	//"fmt"
	//"github.com/google/gopacket"
	//"github.com/google/gopacket/layers"
	//"github.com/google/gopacket/pcap"
	//"github.com/google/gopacket/pcapgo"
	"github.com/steklyanov/trafficSniffer/cmd"
	//"os"
)

//var (
//	deviceName  string = "wlp2s0"
//	snapshotLen int32  = 1024
//	promiscuous bool   = false
//	err         error
//	timeout     time.Duration = -1 * time.Second
//	handle      *pcap.Handle
//	packetCount int = 0
//)

func main() {
	cmd.Execute()
	// Open output pcap file and write header
	//f, _ := os.Create("test.pcap")
	//w := pcapgo.NewWriter(f)
	//w.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)
	//defer f.Close()
	//
	//// Open the device for capturing
	//handle, err = pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	//if err != nil {
	//	fmt.Printf("Error opening device %s: %v", deviceName, err)
	//	os.Exit(1)
	//}
	//defer handle.Close()
	//
	//// Start processing packets
	//packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	//for packet := range packetSource.Packets() {
	//	// Process packet here
	//	fmt.Println(packet)
	//	w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
	//	packetCount++
	//
	//	// Only capture 100 and then stop
	//	if packetCount > 100 {
	//		break
	//	}
	//}
}
