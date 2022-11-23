package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("superkeymode1039JnAmm")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	//client := &http.Client{}
	//req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	//	req.Header.Set("Token", validToken)
	//res, err := client.Do(req)
	//	if err != nil {
	//		fmt.Fprintf(w, "Error: %s", err.Error())
	//	}

	//	body, err := ioutil.ReadAll(res.Body)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Fprintf(w, string(body))

	fmt.Fprintf(w, string(validToken))
}

func blankPage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, string("Welcome to OCI blank screen, use /login to get your token! -- VERSION 2.0"))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Oracle LAD A-Team"
	claims["exp"] = time.Now().Add(time.Minute * 90).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func handleRequests() {

	http.HandleFunc("/", blankPage)
	http.HandleFunc("/login", homePage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
