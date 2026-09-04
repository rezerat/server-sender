package main

import (
	"fmt"
	"goServer/client"
	"io"
	"net"
)


func main() {
    
    cfg, isRunFirstly, err := client.InitConfig() 
    if err != nil {
        return
    }
    conn, err := net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
        fmt.Println("Server is not reachable.")
    }

    defer conn.Close()

    if isRunFirstly {
        fmt.Println("Start registration client")
        client.SendRegisterPacket(conn, cfg)
    } else {
        fmt.Println("Loggin in...")
        client.SendLoginPacket(conn, cfg)
    }
    
    ansBuf := make([]byte, 1)
    _, _ = io.ReadFull(conn, ansBuf)
    if ansBuf[0] == 0x01 {
        return
    }

}