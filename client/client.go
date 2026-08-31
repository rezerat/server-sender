package client

import (
	"net"
	"os"
)


type config struct {
    ClientId string
    PrivateKey []byte
    PublicKey []byte
}

const configName = "client_config.json"

func initConfig() {
    
}

func SendHello() {
    conn, err := net.Dial("tcp", "0.0.0.0:8080")
    if err != nil {
        println(err)
    }

    opCode := []byte{1}
    reg = true
    conn.Write(opCode)
    os.Stat()
    


}