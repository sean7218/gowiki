package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)


func generateJWT() (string, error) {
	hmacSampleSecret := []byte("AllYourBase")

	// Create the claims
	claims := &jwt.MapClaims{
		"username":"sean7218",
		"password":"123",
		"email":"sean872@g.com",
		"standard":jwt.StandardClaims{
			ExpiresAt: 5,
			Issuer: "sean7218",
		},
	}

	// Signing the key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(hmacSampleSecret)
	fmt.Printf("signed string: %v \nError: %v \n", ss, err)
	return ss, err
}

func verifyJWT(bear string) bool {
	hmacSampleSecret := []byte("AllYourBase")
	var result bool
	// Parse and validate the jwt
	token, err := jwt.Parse(bear, func(token *jwt.Token) (interface{}, error){
		// Validate the algorithm is what you expect
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if ok == false {
			return nil, fmt.Errorf("unexpected signing Mmthod: %v \n", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["email"])
		fmt.Println("authJWT has passed")
		result = true
	} else {
		fmt.Println("authJWT is not good")
		fmt.Println(err)
		result = false
	}

	return  result

}

