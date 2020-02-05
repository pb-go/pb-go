package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/pb-go/pb-go/config"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/chacha20poly1305"
	"log"
	"time"
)

func GenBlake2B(data []byte) string {
	hashsum := blake2b.Sum256(data)
	return hex.EncodeToString(hashsum[:])
}

func EncryptData(src []byte, passwd []byte) ([]byte, string, error) {
	// ciphertext output as []byte, original text hash output as string
	// Passwd 32Bytes, Nonce 12Bytes
	// if Passwd is not satisfying 32 Bytes, use Original Default Encryption Key instead.
	var nonce = []byte(config.ServConf.Security.Encryption_nonce)
	var usedpwd []byte
	if len(passwd) != 32 {
		usedpwd = []byte(config.ServConf.Security.Encryption_key)
	} else {
		usedpwd = passwd
	}
	aead, err := chacha20poly1305.New(usedpwd)
	if err != nil {
		log.Fatalln(err)
	}
	hashedpwd := GenBlake2B(passwd)
	ciphertext := aead.Seal(nil, nonce, src, []byte(hashedpwd))
	return ciphertext, hashedpwd, err
}

func DecryptData(src []byte, passwd []byte) ([]byte, error) {
	var nonce = []byte(config.ServConf.Security.Encryption_nonce)
	var usedpwd []byte
	if len(passwd) != 32 {
		usedpwd = []byte(config.ServConf.Security.Encryption_key)
	} else {
		usedpwd = passwd
	}
	aead, err := chacha20poly1305.New(usedpwd)
	if err != nil {
		log.Fatalln(err)
	}
	currentpwdhash := GenBlake2B(passwd)
	var plaintext []byte
	plaintext, err = aead.Open(nil, nonce, src, []byte(currentpwdhash))
	if err != nil {
		log.Println(err)
	}
	return plaintext, err
}

func GetUTCTimeHash(masterkey string) string {
	hmasterkey := "{" + masterkey + "}"
	currentTime := "{" + string(time.Now().UTC().Format(time.RFC822)) + "}"
	data := []byte(hmasterkey+currentTime)
	hashed := fmt.Sprintf("%x", md5.Sum(data))
	return hashed
}

func GetNanoID() (string, error) {
	id, err := gonanoid.Nanoid(4)
	if err != nil {
		log.Fatalln("Failed to generate nanoid!")
	}
	return id, err
}

