package api

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"log"
	"queue-blockchain-seminar/blockchain/dao"
	"queue-blockchain-seminar/blockchain/util"
	"strings"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func ApiServer() {
	r := mux.NewRouter()
	r.HandleFunc("/user", GetUserByPublicKey).Methods("GET")
	r.HandleFunc("/users", GetAllUsers).Methods("GET")
	if err := http.ListenAndServe(":3000", handlers.CORS()(r)); err != nil {
		log.Fatal(err)
	}
}

func GetUserByPublicKey(w http.ResponseWriter, r *http.Request) {
	publickey := r.FormValue("publicKey")
	pubKeyBytes, err := base64.StdEncoding.DecodeString(publickey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	publicKey := strings.ToUpper(util.ByteToHex(pubKeyBytes))
	user, errDb := dao.GetUserByPubKey(publicKey)
	if errDb != nil {
		respondWithError(w, http.StatusInternalServerError, errDb.Error())
		return
	}
	respondWithJson(w, http.StatusOK, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	publickey := r.FormValue("publicKey")
	pubKeyBytes, err := base64.StdEncoding.DecodeString(publickey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	publicKey := strings.ToUpper(util.ByteToHex(pubKeyBytes))
	users, errDb := dao.GetAllUsers(publicKey)
	if errDb != nil {
		respondWithError(w, http.StatusInternalServerError, errDb.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
