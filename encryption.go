package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func enc(key, plainText []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cipherText := make([]byte, len(plainText) + block.BlockSize())

	iv := cipherText[:block.BlockSize()]
	rand.Read(iv)

	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(cipherText[block.BlockSize():], plainText)

	return cipherText, err
}

func dec(key, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainText := make([]byte, len(cipherText) - block.BlockSize())

	iv := cipherText[:block.BlockSize()]

	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(plainText, cipherText[block.BlockSize():])

	return plainText, err
}
