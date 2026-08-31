package main

import (
	"fmt"
	"goServer/server"
	"goServer/storage"
	"net"
)

func main() {
    err := storage.Сonnect()
    if err != nil {
        fmt.Println("Server can't connect to the database %w\n", err)
        return
        
    }
    // connected := true

    fmt.Println("Hello from Club server)")
    channel, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Port is missing...")
        return 
    }

    defer channel.Close()

    for {
        conn, err := channel.Accept()
        if err != nil {
            fmt.Printf("Something went wrong! \n%v", err)
            continue
        }
        go server.HandleConnection(conn)
    }


}