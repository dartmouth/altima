package main

/*
References
- https://www.calhoun.io/creating-random-strings-in-go/
- https://blog.logrocket.com/learn-golang-encryption-decryption/
- https://gobyexample.com/writing-files
- https://stackoverflow.com/questions/13514184/how-can-i-read-a-whole-file-into-a-string-variable/38811255#38811255
*/

/*
Example for testing
$ go run cred/encrypt.go
Encrypted string: 3ZHmEGf88szfbNGPfc9694VJRDfUTQ==
Decrypted string: hello encrypted world!
*/

import (
  "fmt"
  "os"
  "time"
  "math/rand"
  "crypto/aes"
  "crypto/cipher"
  "encoding/base64"
)

const keyFile = "altima_key.pem"
var encryptionSalt = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func GenerateRandomString(length int) string {
  const charset = "abcdefghijklmnopqrstuvwxyz" +
    "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  var seededRand *rand.Rand = rand.New(
    rand.NewSource(time.Now().UnixNano()))
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func CreateKeyFileIfNotExists() {
  _, err := os.Stat(keyFile)
  if os.IsNotExist(err) {
    f, err := os.Create(keyFile)
    check(err)
    f.WriteString(GenerateRandomString(32))
    f.Close()
  }
}

func GetEncryptionKey() string {
  b, err := os.ReadFile(keyFile)
  check(err)
  return string(b)
}

func Encrypt(text, key string) string {
  block, err := aes.NewCipher([]byte(key))
  if err != nil {
    return ""
  }
  plainText := []byte(text)
  cfb := cipher.NewCFBEncrypter(block, encryptionSalt)
  cipherText := make([]byte, len(plainText))
  cfb.XORKeyStream(cipherText, plainText)
  return base64.StdEncoding.EncodeToString(cipherText)
 }

func Decrypt(text, key string) string {
  block, err := aes.NewCipher([]byte(key))
  if err != nil {
    return ""
  }
  cipherText, err := base64.StdEncoding.DecodeString(text)
  if err != nil {
    panic(err)
  }
  cfb := cipher.NewCFBDecrypter(block, encryptionSalt)
  plainText := make([]byte, len(cipherText))
  cfb.XORKeyStream(plainText, cipherText)
  return string(plainText)
}

func main() {
  CreateKeyFileIfNotExists()
  key := GetEncryptionKey()
  encryptedString := Encrypt("hello encrypted world!", key)
  fmt.Println("Encrypted string: " + encryptedString)
  fmt.Println("Decrypted string: " + Decrypt(encryptedString, key))
}
