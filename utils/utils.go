package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"strings"
)

const (
	SECRET_TEXT       = "deltaphone store"
	SECRET_KEY_PREFIX = "Ciph"
)

func GenerateSecretKey(key string) string {

	bReader := bytes.NewReader([]byte(SECRET_TEXT))

	block, err := aes.NewCipher([]byte(fmt.Sprintf("%s%s", SECRET_KEY_PREFIX, key)))
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

func PathFilter(path string) string {
	return strings.Trim(strings.Trim(path, " "), "/")
}
