package main

import (
	"net/http"
	"fmt"

)

// https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81

type Adapter func(http.Handler) http.Handler

func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func isAuthenticated() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

			tok := r.PostFormValue("bearer")
			htok := r.Header.Get("bearer")
			fmt.Println(tok)
			fmt.Println(htok)
			if len(tok) < 10 {
				panic(h)
			}
			fmt.Println("isAuthenticaed Fired!!!!")
			h.ServeHTTP(w, r)
		})
	}
}

func sendProtected() http.Handler {
	return http.HandlerFunc(setupUsers)
}