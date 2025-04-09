package services

type EncrtyptService interface {
	Encrypt(pwd []byte ) (string, error)
	Compare(pwd string, hash []byte) error
}