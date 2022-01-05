package main

import (
	"net"

	"github.com/fagongzi/log"
)

const target = "0.0.0.0:6203"

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)

	if err != nil {
		log.Errorf("Server read fail with error: %v", err)
		return
	}

	log.Infof("Xiaoxiao: Read Buf = [%v]", string(buf))

	_, err = conn.Write(buf)
	if err != nil {
		log.Infof("Server write resp to client failed with error: %v", err)
		return
	}
}

func main() {
	// Init log
	log.InitLog()

	// Server, listen 6203
	tcpaddr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		log.Errorf("Resolve TCP addr failed with error: ", err)
		return
	}

	serverfd, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		log.Infof("Listen TCP addr failed with error: ", err)
		return
	}

	defer serverfd.Close()

	// Handle conn
	for {
		conn, err := serverfd.AcceptTCP()
		if err != nil {
			log.Infof("Server accept faile with error: ", err)
			return
		}

		go handleClient(conn)
	}
}
