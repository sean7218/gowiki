package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"encoding/json"
)

type AuthHandler struct {
	db *sql.DB
}

type LoginHandler struct {
	db *sql.DB
}

func (h *LoginHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		if username != "" && password != "" {

			pwd := h.findUserPassword(username)
			if pwd == password {
				w.WriteHeader(http.StatusAccepted)
				token, err := generateJWT()
				if err != nil {
					fmt.Fprintln(w, "Error with token generation")
					return
				}
				var out = struct { Token string `json:token` } { token}
				jsn, err := json.Marshal(out)
				if err != nil {
					fmt.Fprintln(w, "Error with token because of the marshal")
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsn)
			} else {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("403 - Access Denied."))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request: Missing Fields."))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request: Incorrect Method."))
	}
}



func registerHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		if r.Method == http.MethodPost {
			username := r.PostFormValue("username")
			email := r.PostFormValue("email")
			password := r.PostFormValue("password")
			if username != "" && email != "" && password !="" {

				stmtIns, err := db.Prepare("INSERT INTO t VALUES (?, ?, ?, ?)")
				if err != nil { panic(err.Error())}
				defer stmtIns.Close()
				_, err = stmtIns.Exec( nil , username, email, password)
				if err != nil { panic(err.Error()) }
				fmt.Fprintln(w, "Finished inserting")

			} else {
				fmt.Fprintln(w, "Missing Parameter")
				return
			}
		} else {
			fmt.Fprintln(w, "Invalid Post")
			return
		}
	})
}

func (h *LoginHandler)findUserPassword(name string) string {
	var outUsername string
	var outID int
	var outEmail string
	var outPassword string


	stmtOut, err := h.db.Prepare("SELECT * FROM t WHERE username = ?")
	if err != nil { panic(err.Error()) }
	stmtOut.QueryRow(name).Scan(&outID, &outUsername, &outEmail, &outPassword)

	return outPassword
}

func (a *AuthHandler)findUserByName(username string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var outUsername string
		var outID int
		var outEmail string
		var outPassword string

		fmt.Fprintln(w, r.URL.Path)
		stmtOut, err := a.db.Prepare("SELECT * FROM t WHERE username = ?")
		if err != nil { panic(err.Error()) }
		conNum := a.db.Stats().OpenConnections
		fmt.Fprintf(w, "Connection Number %d \n", conNum)
		stmtOut.QueryRow(username).Scan(&outID, &outUsername, &outEmail, &outPassword)

		fmt.Fprintf(w, "findUserByName: %s \n", outUsername )
	})
}

func (a *AuthHandler)findUserByEmail(email string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		var outID int
		var outUsername string
		var outEmail string
		var outPassword string

		stmtOut2, err := a.db.Prepare("SELECT * FROM t WHERE email = ?")
		if err != nil { panic(err.Error()) }
		stmtOut2.QueryRow(email).Scan(&outID, &outUsername, &outEmail, &outPassword)

		fmt.Fprintf(w, "User Email: %s", outEmail)
	})

}
