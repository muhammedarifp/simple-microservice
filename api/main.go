package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var MySigningKey = []byte("11wdwd")

func ParseJwt(token string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return MySigningKey, nil
	})

	if err != nil {
		return false
	} else {
		return true
	}
}

func HandleRequst() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Values("Token") != nil {
			if r.Header.Values("Token")[0] != "" {
				status := ParseJwt(r.Header.Values("Token")[0])
				if !status {
					fmt.Fprintln(w, "!! ----------- !! Invalid Token Provided !! ----------- !!", status)
					return
				} else {
					fmt.Fprintln(w, "< --------------- > Its Some Secret < --------------- >")
					return
				}
			} else {
				fmt.Fprintln(w, "!! ----------- !! Token Not Provided !! ----------- !!")
			}
		} else {
			fmt.Fprintln(w, "!! ----------- !! Token Not Provided !! ----------- !!")
		}
	})
	http.ListenAndServe(":9000", nil)
}

func main() {
	HandleRequst()
}
