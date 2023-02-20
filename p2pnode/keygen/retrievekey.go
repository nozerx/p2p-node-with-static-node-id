package keygen

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p/core/crypto"
)

func getFileSize() int {
	file, _ := os.Stat("init.bin")
	return int(file.Size())
}

func RetrieveKey() (crypto.PrivKey, error) {
	keyfile, err := os.Open("init.bin")
	if err != nil {
		fmt.Println("Error while opening the init.txt file")
		keyfile.Close()
		return generateKey()
	}
	keyBytes := make([]byte, getFileSize())
	_, err = keyfile.Read(keyBytes)
	if err != nil {
		fmt.Println("Error while retrieving the key from the init.txt file")
		return nil, err
	}
	fmt.Println("Read ", getFileSize(), " bytes")
	keyData := &KeyData{}
	err = json.Unmarshal(keyBytes, keyData)
	if err != nil {
		fmt.Println("Error while unmarshalling the keyData")
		return nil, err
	}
	privKey, err := crypto.UnmarshalPrivateKey(keyData.PrivKey)
	if err != nil {
		fmt.Println("Error during unmarshalling private key")
	}
	return privKey, nil

}
