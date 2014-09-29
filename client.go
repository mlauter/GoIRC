// Package GoIRC is an irc client
package main

import( "net"
        "log"
        "fmt"
)

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
    int, err := c.Write(msg)
    if err != nil {
        log.Fatalf("%d, %v", int, err)
    }

    msg = []byte("USER cLannister 8 * :Cersei Lannister\r\n")
    int, err = c.Write(msg)
    if err != nil {
        log.Fatalf("%d, %v", int, err)
    } else {
        fmt.Println("I connected")
    }

    buffer := make([]byte, 8192)

    for {
        n, err := c.Read(buffer)
        s := string(buffer[:n])
        if err != nil {
            log.Fatalf("%v", err)
        } else {
            println(s)
        }
    }
}


