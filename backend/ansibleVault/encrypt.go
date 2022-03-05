package ansibleVault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
	"log"
	"strings"
)

func Encrypt(content, password string) (result string, err error) {
	salt, err := GenerateRandomBytes(32)
	check(err)
	key1, key2, iv := genKeyInitctr(password, salt)
	cipherText := createCipherText(content, key1, iv)
	combined := combineParts(cipherText, key2, salt)
	vaultText := hex.EncodeToString([]byte(combined))
	result = formatOutput(vaultText)
	return
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func check(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}

func genKeyInitctr(password string, salt []byte) (key1, key2, iv []byte) {
	keyLength := 32
	ivLength := 16
	key := pbkdf2.Key([]byte(password), salt, 10000, 2*keyLength+ivLength, sha256.New)
	key1 = key[:keyLength]
	key2 = key[keyLength : (keyLength)*2]
	iv = key[(keyLength * 2):(keyLength*2 + ivLength)]
	return
}

func createCipherText(content string, key1, iv []byte) []byte {
	bs := aes.BlockSize //16
	padding := bs - len(content)%bs
	if padding == 0 {
		padding = bs
	} // 1<=x<=16
	padChar := rune(padding)
	padArray := make([]byte, padding)
	for i := range padArray {
		padArray[i] = byte(padChar)
	}
	plainText := []byte(content)
	plainText = append(plainText, padArray...) // mod 16 == 0

	aesCipher, err := aes.NewCipher(key1)
	check(err)
	cipherText := make([]byte, len(plainText))

	aesBlock := cipher.NewCTR(aesCipher, iv)
	aesBlock.XORKeyStream(cipherText, plainText)
	return cipherText
}

func combineParts(cipherText, key2, salt []byte) string {
	hmacEncrypt := hmac.New(sha256.New, key2)
	_, err := hmacEncrypt.Write(cipherText)
	check(err)
	hexSalt := hex.EncodeToString(salt)
	hexHmac := hmacEncrypt.Sum(nil)
	hexCipher := hex.EncodeToString(cipherText)
	combined := string(hexSalt) + "\n" + hex.EncodeToString([]byte(hexHmac)) + "\n" + string(hexCipher)
	return combined
}

func formatOutput(vaultText string) string {
	heading := "$ANSIBLE_VAULT"
	version := "1.1"
	cipherName := "AES256"

	headerElements := []string{heading, version, cipherName}
	header := strings.Join(headerElements, ";")

	elements := []string{header}
	for i := 0; i < len(vaultText); i += 80 {
		end := i + 80
		if end > len(vaultText) {
			end = len(vaultText)
		}
		elements = append(elements, vaultText[i:end])
	}
	elements = append(elements, "")

	whole := strings.Join(elements, "\n")
	return whole
}
