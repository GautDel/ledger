package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Encrypt(data string, envKey string) ([]byte, error){
    plaintext := []byte(data)
    key := []byte(envKey) 

    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, 12)
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

    aesGCM, err := cipher.NewGCM(block)
    if err != nil  {
        return nil, err
    }

    encrypted := aesGCM.Seal(nil, nonce, plaintext, nil)

    return append(nonce, encrypted...), nil 
}

func Decrypt(encrypted []byte, envKey string) ([]byte, error) {

    key := []byte(envKey) 

	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Extract the nonce from the encrypted data
	nonce := encrypted[:12]

	// Create a new GCM cipher block using the AES cipher block and the extracted nonce
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Decrypt the data using AES-GCM
	decrypted, err := aesGCM.Open(nil, nonce, encrypted[12:], nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

func EncryptField(ctx *gin.Context, field string) []byte {
	encrypted, err := Encrypt(field, os.Getenv("ENC_KEY"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to encrypt data", "error": err.Error()})
	}

	return encrypted
}

func DecryptField(ctx *gin.Context, field []byte) string {

	decrypted, err := Decrypt(field, os.Getenv("ENC_KEY"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to decrypt data", "error": err.Error()})
	}

	return string(decrypted)
}
