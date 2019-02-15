package dao

import (
	"gopkg.in/mgo.v2/bson"
	"queue-blockchain-seminar/blockchain/infrastructure"
	"queue-blockchain-seminar/blockchain/dto"
)

var user_cl = infrastructure.SetCollection(infrastructure.Users.String())

func InsertNewUser (user dto.User) (dbErr error) {
	dbErr = user_cl.Insert(user)
	return
}

// userのsig確認
func CheckExistenceOfUser(publicKey string) (int) {
	count, _ := user_cl.Find(bson.M{"public_key": publicKey}).Count()
	return count
}

func GetUserByPubKey(publicKey string) (user dto.User, err error) {
	err = user_cl.Find(bson.M{"public_key": publicKey}).One(&user)
	return
}

func GetAllUsers(publicKey string) (users []dto.User, err error) {
	regex := bson.RegEx{publicKey, ""}
	selector := bson.M{"public_key": bson.M{"$not": regex}}
	err = user_cl.Find(selector).All(&users)
	return
}

func UpdateUserCoin(user dto.User) (dbErr error) {
	selector := bson.M{"_id": user.ID}
	update := bson.M{"$set":user}
	dbErr = user_cl.Update(selector, update)
	return
}
