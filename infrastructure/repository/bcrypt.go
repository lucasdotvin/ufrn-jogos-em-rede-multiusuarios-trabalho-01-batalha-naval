package repository

import (
	"golang.org/x/crypto/bcrypt"
	"trabalho-01-batalha-naval/config"
)

type BcryptHashRepository struct {
	cost int
}

func NewBcryptHashRepository(cfg config.Config) *BcryptHashRepository {
	return &BcryptHashRepository{
		cost: cfg.HashCost,
	}
}

func (b *BcryptHashRepository) Hash(plain string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), b.cost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (b *BcryptHashRepository) Compare(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

	if err != nil {
		return false
	}

	return true
}
