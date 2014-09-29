// Package GoIRC is an irc client
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func usrInput(toConn chan []byte) error {
	buffer := make([]byte, 4096)
	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		toConn <- buffer[:n]
	}
}

func connReader(c *net.TCPConn, toConn chan []byte) error {
	buffer := make([]byte, 8192)
	for {
		n, err := c.Read(buffer)
		s := string(buffer[:n])
		if err != nil {
			log.Fatalf("%v", err)
		}
		if strings.HasPrefix(s, "PING") {
			pong := append([]byte("PONG "), s[5:]...)
			pong = append(pong, []byte("\r\n")...)
			toConn <- pong
			fmt.Printf("PONG %v\r\n", string(s[5:]))
		} else {
			println(s)
		}
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "irc.freenode.net:6667")
	if err != nil {
		log.Fatalf("%v", err)
	}

	c, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("%v", err)
	}

	defer c.Close()

	err = c.SetReadBuffer(4096)
	if err != nil {
		log.Fatalf("%v", err)
	}

	msg := []byte("NICK cLannister \r\n")
	_, err = c.Write(msg)
	if err != nil {
		log.Fatalf("%v", err)
	}

	msg = []byte("USER cLannister 8 * :Cersei Lannister\r\n")
	_, err = c.Write(msg)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("I connected")

	toConn := make(chan []byte)
	go usrInput(toConn)
	go connReader(c, toConn)

	for {
		msg := <-toConn
		fmt.Printf("%v", string(msg))
		n, err := c.Write(msg)
		fmt.Printf("%d", n)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

}
