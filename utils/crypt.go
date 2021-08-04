package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Adapted from https://elithrar.github.io/article/generating-secure-random-numbers-crypto-rand/

type CryptUtil struct{}

func init() {
	assertAvailablePRNG()
}

func assertAvailablePRNG() {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Sprintf("crypto/rand is unavailable: Read() failed with %#v", err))
	}
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

func GenerateRandomNumber(n int) (string, error) {
	const letters = "0123456789"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

func GenerateRandomNumberURLSafe(n int) (string, error) {
	b, err := GenerateRandomNumber(n)
	return b, err
}

// GenerateRandomStringURLSafe returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomStringURLSafe(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func (m *CryptUtil) Generate() (generateToken string, err error) {
	// Example: this will give us a 44 byte, base64 encoded output
	token, err := GenerateRandomStringURLSafe(32)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		panic(err)
	}
	//fmt.Println(token)
	return token, err
	// Example: this will give us a 32 byte output
	//token, err = GenerateRandomString(32)
	//if err != nil {
	//	// Serve an appropriately vague error to the
	//	// user, but log the details internally.
	//	panic(err)
	//}
	//fmt.Println(token)
}

func (m *CryptUtil) GenerateNumber(digits int) (generateToken string, err error) {
	// Example: this will give us a 44 byte, base64 encoded output
	number, err := GenerateRandomNumberURLSafe(digits)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		panic(err)
	}
	//fmt.Println(token)
	return number, err
}

