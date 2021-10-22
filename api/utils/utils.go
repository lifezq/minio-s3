package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
)

func GenerateSecretKey(key string) string {

	bReader := bytes.NewReader([]byte("deltaphone minio"))

	block, err := aes.NewCipher([]byte("mini" + key))
	if err != nil {
		panic(err)
	}

	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	var out bytes.Buffer

	writer := &cipher.StreamWriter{S: stream, W: &out}

	if _, err := io.Copy(writer, bReader); err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", out.Bytes())
}
