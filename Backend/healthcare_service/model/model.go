package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pregled struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	PocetakPregleda   int64              `json:"pocetakPregleda" bson:"pocetakPregleda"`
	ZavrsetakPregleda int64              `json:"zavrsetakPregleda" bson:"zavrsetakPregleda"`
	Vakcina           Vakcina            `bson:"vakcina" bson:"vakcina"`
	TipPregleda       TipPregleda        `json:"tipPregleda" bson:"tipPregleda"`
	Gradjanin         *User              `json:"gradjanin" bson:"gradjanin"`
	Lekar             *User              `json:"lekar" bson:"lekar"`
}

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

type Pol string

const (
	Muski  = "Muski"
	Zenski = "Zenski"
)

type Vakcina struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Naziv     string             `json:"naziv" bson:"naziv"`
	Kompanija string             `json:"kompanija" bson:"kompanija"`
}

//const (
//	BCG = "BCG"
//	HB  = "HB"
//	DTP = "DTP"
//	IPV = "IPV"
//	HIB = "HIB"
//	PCV = "PCV"
//)

type TipPregleda string

const (
	Obican      = "Obican"
	Vakcinacija = "Vakcinacija"
)
