package jsonstore

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
	"golang.org/x/crypto/ed25519"
	"queue/seraku-seminar-blockchain/dao"
	"queue/seraku-seminar-blockchain/infrastructure"
	"queue/seraku-seminar-blockchain/static"
	"strings"
)

var _ types.Application = (*JSONStoreApplication)(nil)

// JSONStoreApplication ...
type JSONStoreApplication struct {
	types.BaseApplication
	state state
}

type state struct {
	LastBlockHeight int64
	LastBlockAppHash []byte
	//ValidatorUpdates []types.Validators
}

// NewJSONStoreApplication ...
func NewJSONStoreApplication() *JSONStoreApplication {
	return &JSONStoreApplication{}
}

// byteコードをhexに変換する
func byteToHex(input []byte) string {
	var hexValue string
	for _, v := range input {
		hexValue += fmt.Sprintf("%02x", v)
	}
	return hexValue
}

// Info ...
func (app *JSONStoreApplication) Info(req types.RequestInfo) (resInfo types.ResponseInfo) {
	return types.ResponseInfo{Data: fmt.Sprintf("{\"size\":%v}", 0)}
}

// DeliverTx ... MongoDBの更新を行う
func (app *JSONStoreApplication) DeliverTx(tx []byte) types.ResponseDeliverTx {

	var temp interface{}
	// Unmarshalはデコード先の型が[]byteだった場合にJSONの文字列をBase64エンコーディングとみなして自動的にデコードを行う
	err := json.Unmarshal(tx, &temp)

	if err != nil {
		panic(err)
	}

	message := temp.(map[string]interface{})

	var bodyTemp interface{}

	errBody := json.Unmarshal([]byte(message["body"].(string)), &bodyTemp)
	if errBody != nil {
		panic(errBody)
	}

	body := bodyTemp.(map[string]interface{})

	code := DeliverTx(body, message)

	return types.ResponseDeliverTx{Code: code, Tags: nil}
}

// CheckTx ... Verify the transaction => transactionの有効性確認
func (app *JSONStoreApplication) CheckTx(tx []byte) types.ResponseCheckTx {
	var temp interface{}
	err := json.Unmarshal(tx, &temp)
	if err != nil {
		panic(err)
	}
	message := temp.(map[string]interface{})

	// ==== Signature Validation =======
	pubKeyBytes, err := base64.StdEncoding.DecodeString(message["publicKey"].(string))
	sigBytes, err := hex.DecodeString(message["signature"].(string))
	messageBytes := []byte(message["body"].(string))

	isCorrect := ed25519.Verify(pubKeyBytes, messageBytes, sigBytes)

	if isCorrect != true {
		return types.ResponseCheckTx{Code: static.CodeTypeBadSignature}
	}

	// ==== Signature Validation =======
	var bodyTemp interface{}
	errBody := json.Unmarshal([]byte(message["body"].(string)), &bodyTemp)
	if errBody != nil {
		panic(errBody)
	}
	body := bodyTemp.(map[string]interface{})

	// ==== Does the user really exist? ======
	if body["type"] != "create_user" {
		publicKey := strings.ToUpper(byteToHex(pubKeyBytes))

		// ここでmongoに今のstateでuserが存在しているか確認している
		count := dao.CheckExistenceOfUser(publicKey)

		if count == 0 {
			return types.ResponseCheckTx{Code: static.CodeTypeBadData}
		}
	}

	// ===== Data Validation =======
	codeType := CheckTx(body, message)

	return types.ResponseCheckTx{Code: codeType}
}

// Commit ...Commit the block. Calculate the appHash
func (app *JSONStoreApplication) Commit() types.ResponseCommit {
	appHash := make([]byte, 8)

	count := infrastructure.FindTotalDocuments()

	binary.PutVarint(appHash, count)

	return types.ResponseCommit{Data: appHash}
}

// Query ... Query the blockchain. Unimplemented as of now.
func (app *JSONStoreApplication) Query(reqQuery types.RequestQuery) (resQuery types.ResponseQuery) {
	return
}

// BeginBlock implements the ABCI application interface.
func (app *JSONStoreApplication) BeginBlock(req types.RequestBeginBlock) (res types.ResponseBeginBlock) {
	//app.state.LastBlockHeight = req.Header.Height
	// TODO get Validator address
	fmt.Println("【BeginBlock】", res.String())
	return
}

func (app *JSONStoreApplication) EndBlock(req types.RequestEndBlock) (res types.ResponseEndBlock) {

	// TODO Add Validator logic
	//if len(app.state.ValidatorUpdates) != 0 {
	//	// 新たなValidatorを追加する
	//	res.ValidatorUpdates = app.state.ValidatorUpdates
	//	// stateを削除する
	//	app.state.ValidatorUpdates = []types.ValidatorUpdate{}
	//}
	fmt.Println("【EndBlock】", res.String())
	return
}
