package dto

import "gopkg.in/mgo.v2/bson"

type Transaction struct {
	ID          bson.ObjectId `bson:"_id" json:"_id"`
	From       	string        `bson:"from" json:"from"`		// 送り主のPublicKey
	To	  		string        `bson:"to" json:"to"`			// 送り先のPublicKey
	Amount  	int64         `bson:"amount" json:"amount"`
}
