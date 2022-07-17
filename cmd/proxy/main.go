package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const (
	publishCMD = "PUB"
)

func main() {
	ln, err := net.Listen("tcp", ":443")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := &server{}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("failed to accept: %v", err)
		}
		go s.handleConn(conn)
	}
}

type server struct {
	downstream net.Conn
}

func (s *server) handleConn(conn net.Conn) {
	b := make([]byte, 3)
	n, err := io.ReadAtLeast(conn, b, 3)
	if err != nil {
		log.Fatalf("failed to read: %v", err)
	}
	if n != 3 {
		log.Fatalf("failed to read 3 bytes")
	}
	if string(b) != publishCMD {
		fmt.Println("establishing tunnel between, ", conn.RemoteAddr(), "and", conn.LocalAddr())
		io.Copy(s.downstream, conn)
		go io.Copy(conn, s.downstream)
	} else {
		fmt.Println("adding new downstream", conn.RemoteAddr())
		s.downstream = conn
	}
}
