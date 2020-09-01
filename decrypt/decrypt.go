package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

//Decrypt decrypts the  []byte using key and AES algorithm
func Decrypt(edata []byte, key string) []byte {

	cph, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalln("Decrypt: NewCipher ", err)
	}

	block, err := cipher.NewGCM(cph)
	if err != nil {
		log.Fatalln("Decrypt: NewGCM ", err)
	}

	noncelen := block.NonceSize()
	nonce := edata[:noncelen]
	edata = edata[noncelen:]

	data, err := block.Open(nil, nonce, edata, nil)
	if err != nil {
		log.Fatalln("Decrypt: Open ", err)
	}
	return data
}
