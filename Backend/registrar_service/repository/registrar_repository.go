package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	domain "registrar_service/model/entity"
)

type RegistrarRepository interface {
	GetAllUsers() ([]*domain.User, error)
	GetUserJMBG(jmbg string) (*domain.User, error)
	CreateNewUser(user domain.User) error
	DoctorCreateUser(user *domain.User) error
	ParentCreateUser(user *domain.User) error
	GetNewbornByMotherJMBG(jmbgMajke string) ([]*domain.User, error)
	GetChildrenByParentJMBG(jmbg string, pol string) ([]*domain.User, error)
	IsUserExist(jmbg string) bool
	IsPotvrdaExist(jmbg string) bool
	FindOneUser(jmbg string) *domain.User
	FindOneUserID(id primitive.ObjectID) (*domain.User, error)
	GetChildren(jmbg string, pol domain.Pol) []domain.User
	DeleteUserID(id primitive.ObjectID) error
	DeleteAllUsers() error
	DeleteAllPotvrde() error
	PostPotvrdaSmrti(potvrdaOSmrti domain.PotvrdaSmrti) error
	GetPotvrdaSmrtiJMBG(jmbg string) (*domain.PotvrdaSmrti, error)
	GetAllPotvrdeSmrti() ([]*domain.PotvrdaSmrti, error)
	DeletePotvrdaSmrtiID(id primitive.ObjectID) error
}
