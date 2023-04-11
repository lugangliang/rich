package main

import (
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn, targetAddr string) {
	// 建立到目标服务器的连接
	targetConn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		log.Printf("Failed to connect to %s: %v\n", targetAddr, err)
		return
	}
	defer targetConn.Close()

	log.Printf("Connected to target %s\n", targetAddr)

	// 将源连接和目标连接相互关联
	go func() {
		_, err := io.Copy(targetConn, conn)
		if err != nil {
			log.Printf("Error copying to target: %v\n", err)
		}
	}()
	_, err = io.Copy(conn, targetConn)
	if err != nil {
		log.Printf("Error copying to source: %v\n", err)
	}
}

func main() {
	listenAddr := ":9000"
	targetAddr := "localhost:8000"

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v\n", listenAddr, err)
	}
	defer listener.Close()

	log.Printf("Listening on %s and proxying to %s\n", listenAddr, targetAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v\n", err)
			continue
		}
		go handleConnection(conn, targetAddr)
	}
}
