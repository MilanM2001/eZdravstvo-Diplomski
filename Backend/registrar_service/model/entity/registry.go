package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Ime           string             `json:"ime" bson:"ime"`
	Prezime       string             `json:"prezime" bson:"prezime"`
	JMBGOca       string             `json:"jmbgOca" bson:"jmbgOca"`
	JMBGMajke     string             `json:"jmbgMajke" bson:"jmbgMajke"`
	JMBG          string             `json:"jmbg" bson:"jmbg" unique:"true"`
	Pol           Pol                `json:"pol" bson:"pol"`
	DatumRodjenja int64              `json:"datumRodjenja" bson:"datumRodjenja"`
	MestoRodjenja string             `json:"mestoRodjenja" bson:"mestoRodjenja"`
}

type Pol string

const (
	Muski  = "Muski"
	Zenski = "Zenski"
)

type Credentials struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	JMBG     string             `bson:"jmbg" json:"jmbg" validate:"onlyCharAndNum,required"`
	Password string             `bson:"password" json:"password" validate:"onlyCharAndNum,required"`
	UserType UserType           `bson:"userType" json:"userType" validate:"onlyChar"`
}

type PotvrdaSmrti struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	JMBG       string             `json:"jmbg" bson:"jmbg"`
	DatumSmrti int64              `json:"datumSmrti" bson:"datumSmrti"`
	MestoSmrti string             `json:"mestoSmrti" bson:"mestoSmrti"`
}

type UserType string

const (
	Admin     = "Admin"
	Regular   = "Regular"
	Doctor    = "Doctor"
	Registrar = "Registrar"
)
