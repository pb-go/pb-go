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

// GenBlake2B : Pack Blake2B-256 SUM to hex string
func GenBlake2B(data []byte) string {
	hashsum := blake2b.Sum256(data)
	return hex.EncodeToString(hashsum[:])
}

// EncryptData : ciphertext output as []byte, original text hash output as string
func EncryptData(src []byte, passwd []byte) ([]byte, string, error) {
	// Passwd 32Bytes, Nonce 12Bytes
	// if Passwd is not satisfying 32 Bytes, use Original Default Encryption Key instead.
	var nonce = []byte(config.ServConf.Security.EncryptionNonce)
	var usedpwd []byte
	if len(passwd) != 32 {
		usedpwd = []byte(config.ServConf.Security.EncryptionKey)
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

// DecryptData : Reverse action as for EncryptData
func DecryptData(src []byte, passwd []byte) ([]byte, error) {
	var nonce = []byte(config.ServConf.Security.EncryptionNonce)
	var usedpwd []byte
	if len(passwd) != 32 {
		usedpwd = []byte(config.ServConf.Security.EncryptionKey)
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

// GetUTCTimeHash : MD5 Hash for verification of administration and masterkey
func GetUTCTimeHash(masterkey string) string {
	hmasterkey := "{" + masterkey + "}"
	currentTime := "{" + string(time.Now().UTC().Format(time.RFC822)) + "}"
	data := []byte(hmasterkey + currentTime)
	hashed := fmt.Sprintf("%x", md5.Sum(data))
	return hashed
}

// GetNanoID : Nano ID Getter
func GetNanoID() (string, error) {
	// the nano id length is fixed as 4 chars,
	// BUT you do REMEMBER if you wanna change it,
	// you should change the byte buffer length in VerifyRECAPTCHA Function as well
	// otherwise, you will never get correct res from db.
	id, err := gonanoid.Nanoid(4)
	if err != nil {
		log.Fatalln("Failed to generate nanoid!")
	}
	return id, err
}
