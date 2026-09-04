package client

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)


type config struct {

    ClientId string  `json:"Client-id"`
    PrivateKey []byte `json:"Private-key"`
    PublicKey []byte  `json:"Public-key"`
    IsRegistered bool  `json:"Registration"`

}

const configName = "client_config.json"

func InitConfig() (*config, bool, error){
    data, err := os.ReadFile(getConfigPath())
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
            IsRegistered: false,
        }

        err = saveConfig(&cfg)
        if err != nil {
            return nil, false, fmt.Errorf("Can't generate config for user : %v\n", err)
        }
        return &cfg, true, nil
    }
    return nil, false, err
}

func saveConfig(cfg *config) error {
    fileData, _ := json.MarshalIndent(cfg, "", " ")
    err := os.WriteFile(getConfigPath(), fileData, 0600)
        if err != nil {
            return fmt.Errorf("Can't save config file: %v", err)
        }
    return nil
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


func SendRegisterPacket(conn net.Conn, cfg *config) error {
    var payload []byte
    payload = append(payload, 1)

    idBytes := []byte(cfg.ClientId)
	payload = append(payload, byte(len(idBytes)))

    payload = append(payload, idBytes...)
    payload = append(payload, cfg.PublicKey...)

    _, err := conn.Write(payload)
	if err != nil {
		return fmt.Errorf("Error sending register packet: %v\n", err)
	}

    responsebuf := make([]byte, 1)
    _, err = io.ReadFull(conn, responsebuf)
    if err != nil {
        return fmt.Errorf("No response from the server after registation %v\n", err)
    }
    switch responsebuf[0] {
    case byte(0x01):
        cfg.IsRegistered = true
        saveConfig(cfg)
    case byte(0x02):
        // TODO
        InitConfig()
    }
    return nil
}

func SendLoginPacket(conn net.Conn, cfg *config) error{
    var payload []byte
    payload = append(payload, 2)
    idBytes :=  []byte(cfg.ClientId)
    payload = append(payload, byte(len(idBytes)))
    payload = append(payload, idBytes...)
    _, err := conn.Write(payload)

    if err != nil {
        return fmt.Errorf("Log in error :%v\n", err)
    }
    return nil
}
