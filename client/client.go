package client

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)


type config struct {
    ClientId string
    PrivateKey []byte
    PublicKey []byte
}

const configName = "client_config.json"

func InitConfig() (*config, bool, error){
    configPath := getConfigPath()
    data, err := os.ReadFile(configPath)
    if err == nil {
        var cfg config
        if err := json.Unmarshal(data, &cfg); err != nil {
            return nil, false, err
        }
        return &cfg, false, nil
    }  
    
    if os.IsNotExist(err) {
        fmt.Println("Generate config for new User...")
        newClientId := uuid.New().String()

        pubKey, prKey, err := ed25519.GenerateKey(rand.Reader)
        if err != nil {
            return nil, false, fmt.Errorf("Key generation error: %v", err)
        }
        cfg := config{
            ClientId: newClientId,
            PublicKey: pubKey,
            PrivateKey: prKey,
        }
        fileData, _ := json.MarshalIndent(cfg, "", " ")

        err = os.WriteFile(configPath, fileData, 0600)
        if err != nil {
            return nil, false, fmt.Errorf("Can't save config file: %v", err)
        }
        return &cfg, true, nil
    }
    return nil, false, err
}

func getConfigPath() string {
    configDir, err := os.UserConfigDir()
    if err != nil {
        return configName
    }
    appDir := filepath.Join(configDir, "Server-Client")
    os.MkdirAll(appDir, os.ModePerm)
    return filepath.Join(appDir, configName)
}


func SendRegisterPacket(conn net.Conn, cfg *config){
    var payload []byte
    payload = append(payload, 1)

    idBytes := []byte(cfg.ClientId)
	payload = append(payload, byte(len(idBytes)))

    payload = append(payload, idBytes...)
    payload = append(payload, cfg.PublicKey...)

    _, err := conn.Write(payload)
	if err != nil {
		fmt.Printf("Error sending register packet: %v\n", err)
	}
}

// func SendHello() {
//     conn, err := net.Dial("tcp", "0.0.0.0:8080")
//     if err != nil {
//         println(err)
//     }

//     opCode := []byte{1}
//     reg = true
//     conn.Write(opCode)
//     os.Stat()
    


// }