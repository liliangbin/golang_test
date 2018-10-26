package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/codegangsta/negroni"
)

const (
	SecretKey = "welcome"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type Response struct {
	Data string `json:"data"`
}
type Token struct {
	Token string `json:"token"`
}

func StartServer() {

	http.Handle("/resource", negroni.New(
		negroni.HandlerFunc(ValidateToeknMiddleWare),
		negroni.Wrap(http.HandlerFunc(ProtectHander)),
	))



	http.HandleFunc("/login", LoginHandler)

	log.Println("Now listening...")
	http.ListenAndServe(":8089", nil)
}

func main() {

	StartServer()

}

func ProtectHander(w http.ResponseWriter, r *http.Request) {

	response := Response{"Going to get access the protect resource "}

	JsonResponse(response, w)

}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	log.Println(user.Password + "    fffff")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "error in request")
		return
	}
	log.Println("第二步")
	if strings.ToLower(user.Username) != "liliangbin" {
		if user.Password != "liliangbin" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error login ")
			fmt.Fprintf(w, "Invalid credentials")

			return
		}

	}
	log.Println("第三步")
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = "123456"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("lilili"))
	fmt.Println(tokenString, err)
	log.Println(tokenString + "   " + w.Header().Get("status"))
	response := Token{tokenString}
	JsonResponse(response, w)

}

func ValidateToeknMiddleWare(w http.ResponseWriter, r *http.Request,next http.HandlerFunc) {

	r.ParseForm()//这个地方不会出问题了
	fmt.Println(r.Form)
	var tokenString = r.Form["token"][0]
	//tokenString := ""
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("lilili"), nil
	})

	/*	if err==nil {

			if token.Valid {
				next(w,r)
			}else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w,"token is not valid ")
			}

		}else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w,"unauthorized access to this resource")
		}
	*/

	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims["id"], claims["id"])
		} else {
			fmt.Println(err)
		}

		fmt.Println("You look nice today")
		fmt.Fprintf(w,"index")
		next(w,r)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

}

func JsonResponse(response interface{}, w http.ResponseWriter) {

	log.Println("jsonresponse")
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

