package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Credentials struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	JMBG     string             `bson:"jmbg" json:"jmbg"`
	Password string             `bson:"password" json:"password"`
	UserType UserType           `bson:"userType" json:"userType"`
}

type UserType string

const (
	Admin     = "Admin"
	Regular   = "Regular"
	Doctor    = "Doctor"
	Registrar = "Registrar"
)

type Claims struct {
	UserID    primitive.ObjectID `json:"user_id"`
	JMBG      string             `json:"jmbg"`
	Role      UserType           `json:"userType"`
	ExpiresAt time.Time          `json:"expires_at"`
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

type Ozbiljnost string

const (
	Blaga    = "Blaga"
	Umerena  = "Umerena"
	Ozbiljna = "Ozbiljna"
)
