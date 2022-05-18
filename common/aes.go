package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

func Encrypt(plantText, key []byte) (ciphertext []byte, err error) {
	// 捕捉 panic
	defer func() {
		if info := recover(); info != nil {
			err = errors.New(fmt.Sprintf("Panic: %s", info))
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plantText = PKCS7Padding(plantText, block.BlockSize())
	blockModel := cipher.NewCBCEncrypter(block, key)
	ciphertext = make([]byte, len(plantText))
	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Decrypt(ciphertext, key []byte) (plantText []byte, err error) {
	// 捕捉 panic
	defer func() {
		if info := recover(); info != nil {
			err = errors.New(fmt.Sprintf("Panic: %s", info))
		}
	}()
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes) //选择加密算法
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, keyBytes)
	plantText = make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText, err = PKCS7UnPadding(plantText, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return plantText, nil
}

func PKCS7UnPadding(plantText []byte, blockSize int) ([]byte, error) {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	if unpadding > length {
		return nil, errors.New("PKCS7UnPadding error")
	}
	return plantText[:(length - unpadding)], nil
}
