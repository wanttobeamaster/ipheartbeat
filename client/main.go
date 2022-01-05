package main

import (
	"bytes"
	"io/ioutil"
	"net"
	"os"

	"github.com/fagongzi/log"
)

const target = "127.0.0.1:6203"

func main() {
	// Init log
	log.InitLog()

	// Get Hostname
	hostname, err := os.Hostname()
	if err != nil {
		log.Errorf("Get Hostname failed with error: %v", err)
		return
	}
	log.Infof("Hostname[%v]", hostname)

	// Get locate ip address
	localAddrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Errorf("Get locate ip failed with error: ", err)
		return
	}

	localAddr := ""
	for _, address := range localAddrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && ipnet.IP.String()[:8] == "121.248." {
				localAddr = ipnet.IP.String()
			}

		}
	}
	log.Infof("Locate Ip = [%v]", localAddr)

	// Create tcp connection
	addr, err := net.ResolveTCPAddr("tcp4", target)
	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		log.Errorf("Connect to cloud server failed with error: %v", err)
		return
	}
	defer conn.Close()

	buf := bytes.NewBuffer(nil)
	buf.WriteString(hostname)
	buf.WriteString("\t")
	buf.WriteString(localAddr)

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		log.Errorf("Sent heartbeat to server failed with error: %v", err)
		return
	}

	_, err = ioutil.ReadAll(conn)
	if err != nil {
		log.Errorf("Receive heartbeat response failed with error: %v", err)
		return
	}
	log.Infof("Sent heartbeat succeed!")
}
