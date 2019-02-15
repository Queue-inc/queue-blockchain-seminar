package dto

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID          bson.ObjectId 	`bson:"_id" json:"_id"`
	Name       	string        	`bson:"name" json:"name"`
	PublicKey  	string        	`bson:"public_key" json:"public_key"`
	RemainCoin  int        		`bson:"remain_coin" json:"remain_coin"`
}
