package repository

import (
	domain "auth_service/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRepository interface {
	IsJMBGUnique(jmbg string) bool
	Register(credentials domain.Credentials)
	GetCredentials(jmbg string) (*domain.Credentials, error)
	GetAllCredentials() ([]*domain.Credentials, error)
	DeleteCredentialsID(id primitive.ObjectID) error
}
