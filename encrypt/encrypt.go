package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

//Encrypt encrypts []byte using AES and returns error if any
func Encrypt(data []byte, key string) []byte {

	cph, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalln("Encrypt: ", err)
	}

	block, err := cipher.NewGCM(cph)
	if err != nil {
		log.Fatalln("Encrypt: ", err)
	}

	nonce := make([]byte, block.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		log.Fatalln("Encrypt: ", err)
	}

	return block.Seal(nonce, nonce, data, nil)
}
