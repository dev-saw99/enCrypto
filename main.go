package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dev-saw99/enCrypto/decrypt"

	"github.com/dev-saw99/enCrypto/encrypt"
)

const (
	logo = `
+----------+
| enCrypto |
+----------+
`

	//key needs to be 16byte or 32byte
	key = "thisisthepasskeytoencryptdecrypt"
)

func readFile() []byte {
	bytes, err := ioutil.ReadFile("./data.csv")
	if err != nil {
		log.Fatalln("readFile: ", err)
	}
	return bytes
}

func writeFile(data []byte) {
	err := ioutil.WriteFile("./encrypted.data", data, 0777)
	if err != nil {
		log.Fatalln("writeFile: ", err)
	}
}
func main() {

	fmt.Println(logo)

	bt := readFile()
	fmt.Printf("Encrypting....\n\n")
	ebt := encrypt.Encrypt(bt, key)

	writeFile(ebt)
	fmt.Printf("Encrypted data written to file\n\n")

	fmt.Printf("Decrypting....\n\n")
	dbt := decrypt.Decrypt(ebt, key)

	fmt.Printf("Decrypted data\n\n")
	fmt.Println(string(dbt))

}
