package server

import (
	"fmt"
	"goServer/protoError"
	"io"
	"math/rand"
	"net"
)

var (
    nouns = []string{"Potato", "Cube", "Pan", "Chair", "Banana"}
    adjectives = []string{"Quantum", "Rusty", "Small", "Tall", "Immersive"}
)

type user struct {
    clientId string
    pubKey []byte

}

const (
	opRegister byte = 1
	opLogin    byte = 2
    statusSuccess byte = 0x01
    missingPubKey byte = 0x02
    statusError   byte = 0x02
    opCodeMissing byte = 0x03
    missingClientIdSize byte = 0x5
    missingClientId byte = 0x4
    regErrorCode  byte = 0x99
)

func generateNickname() string {
    noun := nouns[rand.Intn(len(nouns))]
    adj := adjectives[rand.Intn(len(nouns))]
    return fmt.Sprintf("%s-%s", adj, noun)
}

func HandleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Printf("[HC] New connection from %s\n", conn.RemoteAddr().String())

    opCode, clientId, pubKey, err := readClientHello(conn)
    if err != nil {
        fmt.Printf("Packet read error: %v\n", err)
        return
    }
    fmt.Println(clientId, pubKey, opCode)

    

    // user, err := getUser(clientId, pubKey)
    // if err != nil {
    //     // TODO::
    //     // need to know all possible errors, to send diff codes
    //     conn.Write([]byte{statusError})
    //     fmt.Printf("Can not get user : %v\n", err)
    //     return
    // }
    // conn.Write([]byte{statusSuccess})
    // fmt.Println(user)
    
}

func readClientHello(conn net.Conn) (byte, string, []byte, error) {
    opBuf := make([]byte, 1)
	if _, err := io.ReadFull(conn, opBuf); err != nil {
		return 0, "", nil, &protoError.ProtocolError{
            OpCode: opCodeMissing,
            ClientMsg: "Server does't get an operation code from client",
            InternalError: err,
        }
	}
    // 1 || 2
	opCode := opBuf[0]

    lenghtBuf := make([]byte, 1)
    _, err := io.ReadFull(conn, lenghtBuf)
    if err != nil {
        return 0, "", nil, &protoError.ProtocolError{
            OpCode: missingClientIdSize,
            ClientMsg: "Server does't get the client Id size",
            InternalError: err,
        }
    }

    idLenght := int(lenghtBuf[0])
    idBuf := make([]byte, idLenght)
    _, err = io.ReadFull(conn, idBuf)
    if err != nil {
        return 0, "", nil, &protoError.ProtocolError{
            OpCode: missingClientId,
            ClientMsg: "Server does't get the client Id",
            InternalError: err,
        }
    }
    var pubKeyBuf []byte
    if opCode == opRegister {
        pubKeyBuf = make([]byte, 32)
        _, err = io.ReadFull(conn, pubKeyBuf)
        if err != nil {
            return 0, "", nil, &protoError.ProtocolError{
                OpCode: regErrorCode,
                ClientMsg: "Error occurs while server try to register the client",
                InternalError: err,
            }
        }
    }
    
    return opCode, string(idBuf), pubKeyBuf, nil
}


// func getUser(clientId string, pubKey []byte) (user, error) {
//     sPubKey, err := storage.GetPublicKey(clientId)
//     if err != nil {
//         if errors.Is(err, pgx.ErrNoRows) {
//             if len(pubKey) == 0 {
//                 return user{}, &ProtocolError{missingPubKey, "Missing public key"}
//             }
//             fmt.Println("New user - will be added...")

//             nickname := generateNickname()
//             println(nickname)
//             err := storage.SaveDevice(clientId, nickname, pubKey)
//             if err != nil {
//                 return user{}, fmt.Errorf("Registration error %v", err)
                
//             }
//             fmt.Printf("User - %v added to database!\n", nickname)
//              // opReg --> opSuccess(client)
//             return user{
//                 clientId: clientId,
//                 pubKey: pubKey,
//             }, nil

//         } else {
//             return user{}, fmt.Errorf("Can not get public key from database \n")
            
//         }
//     } else {
//         return user{
//         clientId: clientId,
//         pubKey:   sPubKey, 
//     }, nil
//     }
// }


// func handleConnection(conn net.Conn) {
//     // close connection after all functions go through
//     defer conn.Close()
//     // store some bytes from client
//     reader := bufio.NewReader(conn)
//     // bytes of handshake message size
//     bytes := make([]byte, 2)
//     _, err := io.ReadFull(reader, bytes)
//     if err != nil {
//         fmt.Printf("Have no data from client! %s", err)
//         return
//     }

//     // num of handshake message size
//     size := binary.BigEndian.Uint16(bytes)
//     if size > 0 {
//         fmt.Println("Try to decrypt handshake message... \n")
//     } else {
//         error.Error(io.EOF)
//     }
//     //
//     handshakeMsgBuf := make([]byte, size)
//     secretKey := []byte("Homo_server_packager_unpackager_")
//     _, err = io.ReadFull(reader, handshakeMsgBuf)

//     handshakeMsgEnc, err := Decrypt(handshakeMsgBuf, secretKey)
//     if err != nil {
//         fmt.Printf("Decrypting error: %v\n", err)
//     }

//     if handshakeMsgBuf уджек

// }

// func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
//     block, err := aes.NewCipher(key)
//     if err != nil {
//         return nil, err
//     }
//     aesGCM, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return nil, err
// 	}

//     nonceSize := aesGCM.NonceSize()

//     if len(ciphertext) < nonceSize {
//         return nil, fmt.Errorf("ciphertext too short")
//     }

//     nonce, encMessage := ciphertext[:nonceSize], ciphertext[nonceSize:]

//     plaintext, err := aesGCM.Open(nil, nonce, encMessage, nil)

//     if err != nil {
//         return nil, err
//     }
//     println("heres")
//     return plaintext, nil
// }