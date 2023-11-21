package model

type AddPregled struct {
	PocetakPregleda   int64       `json:"pocetakPregleda" bson:"pocetakPregleda"`
	ZavrsetakPregleda int64       `json:"zavrsetakPregleda" bson:"zavrsetakPregleda"`
	VakcinaID         string      `json:"vakcinaID" bson:"vakcinaID"`
	TipPregleda       TipPregleda `json:"tipPregleda" bson:"tipPregleda"`
}
