package DataStructures

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Opravilo struct {
	Id                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Naslov             string             `json:"naslov"`
	Opis               string             `json:"opis"`
	DatumDodajanja     time.Time          `json:"datum_dodajanja"`
	PredvidenDatumDela time.Time          `json:"predviden_datum_dela"`
}

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
