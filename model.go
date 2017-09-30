package main

import (
	"fmt"
	//"os"
	"encoding/json"
	"net/http"
	"html/template"
	"log"
)

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type WebClient struct {
	Id int
	Username string
	Email string
	Password string
}
func setupJSON(){

	// Following are the encoding from go into json
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB	))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Following are the decoding from json to go

}

func sendJSON(w http.ResponseWriter, r *http.Request) {

	aUser := WebClient{1, "sean7218", "sean7218@l.com", "123"}

	b, err := json.Marshal(aUser)
	if err != nil {
		fmt.Println("Error for marshal the json")
		return
	}
	fmt.Fprintf(w, "\n\n JSON: %v",string(b))
}

func setupDwg(w http.ResponseWriter, r *http.Request){

	aUser := WebClient{1, "sean7218", "sean7218@l.com", "211"}
	bUser := WebClient{2, "josh7218", "josh7218@l.com", "12323"}
	cUser := WebClient{3, "mike7218", "mike7218@l.com", "121231233"}
	dUser := WebClient{4, "jesse7218", "jesse7218@l.com", "112332123"}
	eUser := WebClient{5, "john7218", "john7218@l.com", "121231231231233"}
	fUser := WebClient{6, "eric7218", "eric7218@l.com", "121231231231231233"}
	gUser := WebClient{7, "jeremy7218", "jeremy7218@l.com", "1412412412412412423"}


	users := [7]WebClient{ aUser, bUser, cUser, dUser, eUser, fUser, gUser}

	for i := range users {
		log.Printf("Key: %s, Value: %v \n", i)
	}

	tmpl := template.Must(template.ParseFiles("public/dwg.html", "public/header.html", "public/footer.html", "public/script.html"))
	err := tmpl.ExecuteTemplate(w, "dwg.html", users)
	if err != nil {
		fmt.Println("Error for executing the template")
		fmt.Println(err.Error())
		return
	}
}