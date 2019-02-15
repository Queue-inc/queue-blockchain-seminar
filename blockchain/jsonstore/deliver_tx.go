package jsonstore

import (
	"encoding/base64"
	"gopkg.in/mgo.v2/bson"
	"queue-blockchain-seminar/blockchain/dto"
	"queue-blockchain-seminar/blockchain/dao"
	"queue-blockchain-seminar/blockchain/static"
	"strconv"
	"strings"
)

//DeliverTx
func DeliverTx(body map[string]interface{}, message map[string]interface{}) (code uint32) {
	code = static.CodeTypeOK

	entity := body["entity"].(map[string]interface{})
	if len(entity) == 0 {
		code = static.CodeTypeBadData
		return
	}

	pubKeyBytes, _ := base64.StdEncoding.DecodeString(message["publicKey"].(string))
	publicKey := strings.ToUpper(byteToHex(pubKeyBytes))

	switch body["type"] {
	case "create_user": // user登録
		code = createUserDeliver(entity, publicKey)
		break
	case "send_funds": // コインの送金
		code = sendFundsDeliver(entity, publicKey)
		break
	}
	return
}

func createUserDeliver(entity map[string]interface{}, publicKey string) (code uint32) {
	// ユーザーが存在しないことを確認
	count := dao.CheckExistenceOfUser(publicKey)
	if count != 0 {
		code = static.CodeTypeBadData
		return
	}

	// insert
	user := dto.User{
		ID: bson.ObjectIdHex(entity["_id"].(string)),
		PublicKey: publicKey,
		Name: entity["name"].(string),
		RemainCoin: 100,
	}
	dao.InsertNewUser(user)
	return
}

func sendFundsDeliver(entity map[string]interface{}, publicKey string) (code uint32) {
	from, _ := dao.GetUserByPubKey(publicKey)
	amount, _ := strconv.Atoi(entity["amount"].(string))

	// 自分の残金が送金額よりもあることを確認
	if from.RemainCoin < amount {
		code = static.CodeTypeBadData
		return
	}

	// 送金先のユーザーが存在することを確認
	to, dbErr := dao.GetUserByPubKey(entity["to"].(string))
	if dbErr != nil {
		code = static.CodeTypeBadData
		return
	}

	// update
	from.RemainCoin -= amount
	to.RemainCoin += amount

	dao.UpdateUserCoin(from)
	dao.UpdateUserCoin(to)

	code = static.CodeTypeOK
	return
}