package Encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

func DecryptFile(password string, FilePath string) error {
	OldFilePath := FilePath[:(len(FilePath) - 4)]
	hashedKey := CreateHash(password)
	ciphertext, err := ioutil.ReadFile(FilePath)
	if err != nil {
		return err
	}
	if len(ciphertext) < aes.BlockSize {
		return errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	encKey := ciphertext[:32+aes.BlockSize][aes.BlockSize:]
	ciphertext = ciphertext[aes.BlockSize+32:]

	block, err := CreateKey(hashedKey)

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encKey, encKey)

	for i := range encKey {
		if encKey[i] != hashedKey[i] {
			return errors.New("Incorrect Password")
		}
	}

	stream.XORKeyStream(ciphertext, ciphertext)

	plaintextFile, err := os.Create(OldFilePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(plaintextFile, bytes.NewReader(ciphertext))
	if err != nil {
		return err
	}
	os.Remove(FilePath)
	return nil
}
