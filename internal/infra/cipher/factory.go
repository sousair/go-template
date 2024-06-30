package cipher

import "golang.org/x/crypto/bcrypt"

type Cipher interface {
	Encrypt(plainText string) (string, error)
	Compare(plainText, cipherText string) (bool, error)
}

type BcryptCipher struct {
	hashCost int
}

var _ Cipher = (*BcryptCipher)(nil)

func NewBcryptCipher(hashCost int) *BcryptCipher {
	return &BcryptCipher{hashCost: hashCost}
}

func (c *BcryptCipher) Encrypt(plainText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), c.hashCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (c *BcryptCipher) Compare(plainText, cipherText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(cipherText), []byte(plainText))
	if err != nil {
		return false, err
	}

	return true, nil
}
