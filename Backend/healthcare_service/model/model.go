package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type Pregled struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	PocetakPregleda   int64              `json:"pocetakPregleda" bson:"pocetakPregleda"`
	ZavrsetakPregleda int64              `json:"zavrsetakPregleda" bson:"zavrsetakPregleda"`
	Vakcina           *Vakcina           `json:"vakcina" bson:"vakcina"`
	TipPregleda       TipPregleda        `json:"tipPregleda" bson:"tipPregleda"`
	Gradjanin         *User              `json:"gradjanin" bson:"gradjanin"`
	Lekar             *User              `json:"lekar" bson:"lekar"`
}

type Vakcina struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Naziv     string             `json:"naziv" bson:"naziv"`
	Kompanija string             `json:"kompanija" bson:"kompanija"`
}

type Alergija struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Naziv string             `json:"naziv" bson:"naziv"`
}

type Invaliditet struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Naziv string             `json:"naziv" bson:"naziv"`
}

type Karton struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	JMBG         string             `json:"jmbg" bson:"jmbg"`
	Alergije     []Alergija         `json:"alergije" bson:"alergije"`
	Invaliditeti []Invaliditet      `json:"invaliditeti" bson:"invaliditeti"`
}

type PotvrdaSmrti struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	JMBG       string             `json:"jmbg" bson:"jmbg"`
	DatumSmrti int64              `json:"datumSmrti" bson:"datumSmrti"`
	MestoSmrti string             `json:"mestoSmrti" bson:"mestoSmrti"`
}

type Pol string

const (
	Muski  = "Muski"
	Zenski = "Zenski"
)

type TipPregleda string

const (
	Obican      = "Obican"
	Vakcinacija = "Vakcinacija"
)

type Ozbiljnost string

const (
	Blaga    = "Blaga"
	Umerena  = "Umerena"
	Ozbiljna = "Ozbiljna"
)
