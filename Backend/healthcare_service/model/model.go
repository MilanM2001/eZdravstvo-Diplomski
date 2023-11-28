package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Ime           string             `json:"ime" bson:"ime"`
	Prezime       string             `json:"prezime" bson:"prezime"`
	ImeOca        string             `json:"ime_oca" bson:"imeOca"`
	JMBGOca       string             `json:"jmbg_oca" bson:"JMBGOca"`
	ImeMajke      string             `json:"ime_majke" bson:"imeMajke"`
	JMBGMajke     string             `json:"jmbg_majke" bson:"JMBGMajke"`
	DatumRodjenja int64              `json:"datum_rodjenja" bson:"datumRodjenja"`
	MestoRodjenja string             `json:"mesto_rodjenja" bson:"mestoRodjenja"`
	JMBG          string             `json:"jmbg" bson:"JMBG" unique:"true"`
	Pol           Pol                `json:"pol" bson:"pol"`
	Preminuo      bool               `json:"preminuo" bson:"Preminuo"`
	DatimSmrti    int64              `json:"datim_smrti" bson:"DatimSmrti"`
	MestoSmrti    string             `json:"mesto_smrti" bson:"MestoSmrti"`
	Drzava        string             `json:"drzava" bson:"Drzava"`
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
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Naziv      string             `json:"naziv" bson:"naziv"`
	Ozbiljnost Ozbiljnost         `json:"ozbiljnost" bson:"ozbiljnost"`
}

type Invaliditet struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Naziv      string             `json:"naziv" bson:"naziv"`
	Opis       string             `json:"opis" bson:"opis"`
	Ozbiljnost Ozbiljnost         `json:"ozbiljnost" bson:"ozbiljnost"`
}

type Karton struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	JMBG         string             `json:"jmbg" bson:"jmbg"`
	Alergije     []Alergija         `json:"alergije" bson:"alergije"`
	Invaliditeti []Invaliditet      `json:"invaliditeti" bson:"invaliditeti"`
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
