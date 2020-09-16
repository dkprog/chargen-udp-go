package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	port := flag.Int("port", 3019, "CHARGEN udp port")
	flag.Parse()
	address := fmt.Sprintf(":%d", *port)
	conn, err := net.ListenPacket("udp", address)
	var start uint = 0
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("Listening for UDP packets on %s", conn.LocalAddr().String())
	for {
		_, addr, err := conn.ReadFrom(nil)
		if err != nil {
			continue
		}
		log.Printf("Received a packet from %s", addr)
		stream := getASCIIStream(start)
		start++
		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		conn.WriteTo(stream, addr)
	}
}

func getASCIIStream(start uint) []byte {
	ascii := getASCIITable()
	var stream [72]byte
	for i := 0; i < len(stream); i++ {
		j := (start + uint(i)) % uint(len(ascii))
		stream[i] = ascii[j]
	}
	return stream[:]
}

func getASCIITable() []byte {
	var len int = 127 - 32
	var ascii []byte = make([]byte, len)
	for i := 0; i < len; i++ {
		ascii[i] = byte(32 + i)
	}
	return ascii
}
