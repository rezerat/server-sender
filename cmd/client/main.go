package main

import (
	"fmt"
	"goServer/client"
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
    }
}