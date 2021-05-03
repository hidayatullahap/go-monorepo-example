package pkg

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Crypt struct {
	hash string
}

func (p *Crypt) HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (p *Crypt) ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func NewCrypt(hash string) *Crypt {
	return &Crypt{hash: hash}
}
