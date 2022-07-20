package core

import (
	"fmt"
	"log"
	"net"
)

func TcpSendMsg(addr, msg string, times int) error {
	log.Printf("Start sending tcp packets to %s", addr)
	tcpaddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}
	socket, err := net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		return err
	}
	defer socket.Close()
	fmt.Println(socket, err)
	for i := 0; i < times; i++ {
		log.Printf("Send msg : %s", msg)
		_, err = socket.Write([]byte(msg + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}
