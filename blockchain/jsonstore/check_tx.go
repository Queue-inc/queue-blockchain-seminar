package jsonstore

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"queue-blockchain-seminar/blockchain/static"
	"strconv"
	"strings"
)

func CheckTx(body map[string]interface{}, message map[string]interface{}) (code uint32) {
	code = static.CodeTypeBadData

	entity := body["entity"].(map[string]interface{})
	if len(entity) == 0 {
		code = static.CodeTypeBadData
		return
	}

	switch body["type"] {
	case "create_user": // user登録
		code = createUserCheck(entity)
		break
	case "send_funds": // コインの送金
		code = sendFundsCheck(entity)
		break
	}
	return
}

func createUserCheck(entity map[string]interface{}) (code uint32){
	if (entity["_id"] == nil) || !bson.IsObjectIdHex(entity["_id"].(string)) {
		code = static.CodeTypeBadData
		return
	}

	if entity["name"] == nil || strings.TrimSpace(entity["name"].(string)) == "" {
		code = static.CodeTypeBadData
		return
	}

	code = static.CodeTypeOK
	return
}

func sendFundsCheck(entity map[string]interface{}) (code uint32) {
	fmt.Println(entity)
	if entity["to"] == nil || strings.TrimSpace(entity["to"].(string)) == "" {
		code = static.CodeTypeBadData
		return
	}

	amount, err := strconv.Atoi(entity["amount"].(string))
	if err != nil {
		code = static.CodeTypeBadData
		return
	}
	if amount < 0 {
		code = static.CodeTypeBadData
		return
	}

	code = static.CodeTypeOK
	return
}