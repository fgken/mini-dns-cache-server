package main

import (
    "fmt"
    "net"
    "bytes"
    "encoding/binary"
)

type DNSHeader struct {
    id      uint16
    flag    uint16
    qdcount uint16
    ancount uint16
    nscount uint16
    arcount uint16
}

func main() {
    ServerAddr, _ := net.ResolveUDPAddr("udp", ":5000")

    ServerConn, _ := net.ListenUDP("udp", ServerAddr)

    defer ServerConn.Close()

    buf := make([]byte, 1024)

    for {
        n, addr, err := ServerConn.ReadFromUDP(buf)
        r := bytes.NewReader(buf)
        dns := DNSHeader{}
        binary.Read(r, binary.BigEndian, &dns.id)
        binary.Read(r, binary.BigEndian, &dns.flag)
        binary.Read(r, binary.BigEndian, &dns.qdcount)
        fmt.Println("Received: ", string(buf[0:n]), " from ", addr)
        fmt.Printf("ID = %#x\n", dns.id)
        fmt.Printf("FLAG = %b\n", dns.flag)
        fmt.Printf("QDCOUNT = %#x\n", dns.qdcount)
        if err != nil {
            fmt.Println("Error: ", err)
        }
    }
}
