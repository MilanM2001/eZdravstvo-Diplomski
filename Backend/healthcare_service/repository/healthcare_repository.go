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
	GetSviSlobodniPregledi() ([]*model.Pregled, error)
	GetPregledID(id primitive.ObjectID) (*model.Pregled, error)
	PostPregled(appointment *model.Pregled) error
	PutPregled(appointment *model.Pregled) error
	DeletePregledID(id primitive.ObjectID) error
	//
	//GetAllVaccinations() ([]*model.Vaccination, error)
	//GetMyVaccinationsDoctor(id primitive.ObjectID) ([]*model.Vaccination, error)
	//GetMyAvailableVaccinationsDoctor(id primitive.ObjectID) ([]*model.Vaccination, error)
	//GetMyTakenVaccinationsDoctor(id primitive.ObjectID) ([]*model.Vaccination, error)
	//GetAllAvailableVaccinations() ([]*model.Vaccination, error)
	//GetMyTakenVaccinationsRegular(id primitive.ObjectID) ([]*model.Vaccination, error)
	//GetVaccinationByID(id primitive.ObjectID) (*model.Vaccination, error)
	//CreateNewVaccination(vaccination *model.Vaccination) error
	//SetVaccination(vaccination *model.Vaccination) error
	//DeleteVaccinationByID(id primitive.ObjectID) error
	//
	//GetAllZdravstvenoStanje() ([]*model.ZdravstvenoStanje, error)
	//GetZdravstvenoStanjeByID(id primitive.ObjectID) (*model.ZdravstvenoStanje, error)
	//GetZdravstvenoStanjeByJMBG(jmbg string) (*model.ZdravstvenoStanje, error)
	//CreateNewZdravstvenoStanje(zdravstvenoStanje *model.ZdravstvenoStanje) error
	//DeleteZdravstvenoStanjeByJMBG(jmbg string) error
}
