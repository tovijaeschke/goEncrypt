package Encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"os"
	//"encoding/hex"
)

func CreateHash(key string) []byte {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(key))
	return h.Sum(nil)
}

func CreateKey(hashedKey []byte) (cipher.Block, error) {
	block, err := aes.NewCipher(hashedKey)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func SecureDelete(FilePath string) error {
	file, err := os.OpenFile(FilePath, os.O_RDWR, 0666)
	defer file.Close()
	// Find out how large is the target file
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	var size int64 = fileInfo.Size()
	// Create byte array filled with zero's
	zeroBytes := make([]byte, size)
	_, err = file.Write([]byte(zeroBytes))
	if err != nil {
		return err
	}
	err = os.Remove(FilePath)
	if err != nil {
		return err
	}
	return nil
}
