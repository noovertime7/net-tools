package core

import (
	"log"
	"net"
)

func UdpSendMsg(addr, msg string, times int) error {
	log.Printf("Start sending udp packets to %s", addr)
	udpaddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	socket, err := net.DialUDP("udp", nil, udpaddr)
	defer socket.Close()
	if err != nil {
		return err
	}
	for i := 0; i < times; i++ {
		log.Printf("Send msg : %s", msg)
		_, err = socket.Write([]byte(msg + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}
