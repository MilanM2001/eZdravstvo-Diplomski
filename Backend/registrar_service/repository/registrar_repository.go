package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	domain "registrar_service/model/entity"
)

type RegistrarRepository interface {
	GetAllUsers() ([]*domain.User, error)
	GetUserJMBG(jmbg string) (*domain.User, error)
	CreateNewBirthCertificate(user domain.User) error
	DoctorCreateUser(user domain.User) error
	GetNewbornByMotherJMBG(jmbgMajke string) ([]*domain.User, error)
	IsUserExist(jmbg string) bool
	FindOneUser(jmbg string) *domain.User
	FindOneUserID(id primitive.ObjectID) (*domain.User, error)
	CreateNewMarriage(marriage domain.ExcerptFromTheMarriageRegister)
	UpdateCertificate(user domain.User) error
	GetChildren(jmbg string, pol domain.Pol) []domain.User
}
