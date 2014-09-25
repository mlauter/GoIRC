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


    err := c.SetReadBuffer(4096)
    if err != nil {
        log.Fatalf("%v", err)
    }

    c.
    
    buffer := make([]byte, 8192)
    n, err := c.Read(buffer)
    if err != nil {
        log.Fatalf("%v", err)
    } else {
        println(n)
    }
}


