package DataStructures

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Opravilo struct {
	Id                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Naslov             string             `json:"naslov"`
	Opis               string             `json:"opis"`
	DatumDodajanja     time.Time          `json:"datum_dodajanja"`
	PredvidenDatumDela time.Time          `json:"previden_datum_dela"`
}
