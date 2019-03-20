package Encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
)

func EncryptFile(password string, FilePath string) error {
	hashedKey := CreateHash(password)
	plaintext, err := ioutil.ReadFile(FilePath)
	if err != nil {
		return nil
	}
	ciphertext := make([]byte, aes.BlockSize+len(hashedKey)+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}
	block, err := CreateKey(hashedKey)
	if err != nil {
		return err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(hashedKey))
	stream.XORKeyStream(ciphertext[aes.BlockSize+len([]byte(hashedKey)):], plaintext)

	// open output file
	encryptedFile, err := os.Create(FilePath + ".enc")
	if err != nil {
		return nil
	}
	defer func() {
		encryptedFile.Close()
		SecureDelete(FilePath)
	}()
	_, err = io.Copy(encryptedFile, bytes.NewReader(ciphertext))
	if err != nil {
		return err
	}
	return nil
}
