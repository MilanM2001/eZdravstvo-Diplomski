package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"healthcare_service/model"
)

type HealthcareRepository interface {
	GetSviPregledi() ([]*model.Pregled, error)
	GetMojiPreglediLekar(id primitive.ObjectID) ([]*model.Pregled, error)
	GetMojiSlobodniPreglediLekar(id primitive.ObjectID) ([]*model.Pregled, error)
	GetMojiZauzetiPreglediLekar(id primitive.ObjectID) ([]*model.Pregled, error)
	GetMojiPreglediGradjanin(id primitive.ObjectID) ([]*model.Pregled, error)
	GetSviSlobodniPregledi() ([]*model.Pregled, error)
	GetPreglediByGradjaninID(id primitive.ObjectID) ([]*model.Pregled, error)
	GetPregledID(id primitive.ObjectID) (*model.Pregled, error)
	PostPregled(appointment *model.Pregled) error
	PutPregled(appointment *model.Pregled) error
	DeletePregledID(id primitive.ObjectID) error

	GetSveVakcine() ([]*model.Vakcina, error)
	GetVakcinaID(id primitive.ObjectID) (*model.Vakcina, error)
	GetVakcinaNaziv(naziv string) (*model.Vakcina, error)
	PostVakcina(tipVakcine *model.Vakcina) error
	PutVakcina(tipVakcine *model.Vakcina) error
	DeleteVakcinaID(id primitive.ObjectID) error

	GetSveAlergije() ([]*model.Alergija, error)
	GetAlergijaID(id primitive.ObjectID) (*model.Alergija, error)
	GetAlergijaNaziv(naziv string) (*model.Alergija, error)
	PostAlergija(alergija *model.Alergija) error
	DeleteAlergijaID(id primitive.ObjectID) error

	GetSveInvaliditete() ([]*model.Invaliditet, error)
	GetInvaliditetID(id primitive.ObjectID) (*model.Invaliditet, error)
	GetInvaliditetNaziv(naziv string) (*model.Invaliditet, error)
	PostInvaliditet(*model.Invaliditet) error
	DeleteInvaliditetID(id primitive.ObjectID) error

	GetSveKartone() ([]*model.Karton, error)
	GetKartoneJMBG(jmbg string) ([]*model.Karton, error)
	GetKartonJMBG(jmg string) (*model.Karton, error)
	PostKarton(karton model.Karton) error
	PutKarton(karton *model.Karton) error
	DeleteKartonID(id primitive.ObjectID) error
	DeleteKartonJMBG(jmbg string) error
	DeleteAllKarton() error
}
