package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)


func generateJWT() (string, error) {
	mySigningKey := []byte("AllYourBase")

	// Create the claims
	claims := &jwt.MapClaims{
		"username":"sean7218",
		"password":"123",
		"email":"sean872@g.com",
		"standard":jwt.StandardClaims{
			ExpiresAt: 2000,
			Issuer: "sean7218",
		},
	}

	// Signing the key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("signed string: %v \nError: %v \n", ss, err)
	return ss, err
}

func verifyJWT(token string) error {
	mySigningKey := []byte("AllYourBase")

	// Parse and validate the jwt
	ptoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error){
		// Validate the algorithm is what you expect
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if ok != false {
			return nil, fmt.Errorf("unexpected signing Mmthod: %v \n", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	return err
}

func setupCJWT(){

	mySigningKey := []byte("AllYourBase")

	// Create the claims
	claims := &jwt.MapClaims{
		"username":"sean7218",
		"password":"123",
		"email":"sean872@g.com",
		"standard":jwt.StandardClaims{
			ExpiresAt: 2000,
			Issuer: "sean7218",
		},
	}

	// Signing the key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("signed string: %v \nError: %v \n", ss, err)


	// Parse and validate the jwt
	ptoken, err := jwt.Parse(ss, func(token *jwt.Token) (interface{}, error){
		// Validate the algorithm is what you expect
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if ok != false {
			return nil, fmt.Errorf("unexpected signing Mmthod: %v \n", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if true {
		fmt.Printf("\n The token valid: %v", ptoken.Valid)
		fmt.Printf("\n The token header: %v", ptoken.Header)
		fmt.Printf("\n The token method: %v", ptoken.Method)
		fmt.Printf("\n The token claims: %v", ptoken.Claims)
	} else {
		fmt.Println("Sorry the token is incorrect")
		fmt.Printf("\n \n \n The valid token: %v", ptoken.Valid)
	}

}

