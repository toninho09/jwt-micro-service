package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "cert/app.rsa"
	pubKeyPath  = "cert/app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

type Response struct {
	Data string `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}

type AuthToken struct{
	Sub string `json:"sub"` //(subject) = Entidade à quem o token pertence, normalmente o ID do usuário;
	Iss  string `json:"iss"` //(issuer) = Emissor do token;
	Exp  string `json:"exp"` //(expiration) = Timestamp de quando o token irá expirar;
	Iat  string `json:"iat"` //(issued at) = Timestamp de quando o token foi criado;
	Aud  string `json:"aud"` //(audience) = Destinatário do token, representa a aplicação que irá usá-lo.
	Data interface{} `json:"data"`
}

func StartServer() {

	http.HandleFunc("/create", CreateHandler)
	http.HandleFunc("/verify", VerifyHandler)
	log.Println("Now listening...")
	http.ListenAndServe(":80", nil)
}

func main() {

	initKeys()
	StartServer()
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {

	var tokenReceive Token
	err := json.NewDecoder(r.Body).Decode(&tokenReceive)
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token,err := jwt.Parse(tokenReceive.Token, func(token *jwt.Token) (interface{}, error) {
		return verifyKey,nil
	})
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	JsonResponse(token.Claims,w)

}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	var createData AuthToken
	err := json.NewDecoder(r.Body).Decode(&createData)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}


	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)

	if createData.Sub != "" {
		claims["sub"] = createData.Sub
	}

	if createData.Iss != "" {
		claims["iss"] = createData.Iss
	}

	if createData.Exp != "" {
		claims["exp"] = createData.Exp
	}else{
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	}

	if createData.Iat != "" {
		claims["iat"] = createData.Iat
	}else{
		claims["iat"] = time.Now().Unix()
	}

	if createData.Aud != "" {
		claims["aud"] = createData.Aud
	}
	claims["data"] = createData.Data

	token.Claims = claims

	tokenString, err := token.SignedString(signKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		fatal(err)
	}

	response := Token{tokenString}
	JsonResponse(response, w)

}

func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}