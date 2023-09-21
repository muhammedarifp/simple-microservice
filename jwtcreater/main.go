package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var MySigningKey = []byte("11wdwd")

func GetJwt() (string, error) {
	claims := jwt.MapClaims{
		"authorized": true, // Changed "authoraized" to "authorized"
		"user_name":  "arif",
		"exp":        time.Now().Add(time.Hour * 48).Unix(), // Corrected the time calculation
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Changed to HS256 for simplicity
	tokenStr, err := token.SignedString(MySigningKey)
	if err != nil {
		log.Fatal("Token Creation error | ", err.Error())
		return "", err
	}
	return tokenStr, nil
}

func HandleRequsts() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		validToken, err := GetJwt()
		if err != nil {
			log.Fatal("Jwt Not Created . ", err.Error())
		} else {
			fmt.Fprintln(w, validToken)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Viper file read filed . ", err.Error())
	}

	fmt.Println("Starting")
	HandleRequsts()
}
