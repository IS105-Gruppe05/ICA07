package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {

	_, err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("message recieved from %v \n %s ", remoteaddr, p)

		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue

		}

		key := []byte("AES256Key-32Characters1234567890")
		ciphertext, _ := hex.DecodeString("965112884118a19f8b4f6483027ef84daa50d526973dfd84ef2200b5e9")

		nonce, _ := hex.DecodeString("ab35ca68732c41b0b5e934ed")

		block, err := aes.NewCipher(key)
		if err != nil {
			panic(err.Error())
		}

		aesgcm, err := cipher.NewGCM(block)
		if err != nil {
			panic(err.Error())
		}

		plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%s\n", plaintext)
		// Output: exampleplaintext

		go sendResponse(ser, remoteaddr)
	}

}
