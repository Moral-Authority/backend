package handlers

import (
	"crypto/sha512"
	"encoding/base64"
	"math/rand"
)

func generateRandomSalt(saltSize int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, saltSize)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func hashPassword(password string, salt string) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)
	// Create sha-512 hasher
	var sha512Hasher = sha512.New()
	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)
	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)
	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	// Convert the hashed password to a base64 encoded string
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	return base64EncodedPasswordHash
}
