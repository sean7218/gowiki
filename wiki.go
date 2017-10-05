package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
	"regexp"
	"errors"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	Title string
	Body  []byte
}

type User struct {
	UserName string
	Email string
	Password string
}

var templates = template.Must(template.ParseFiles("public/edit.html", "public/view.html", "public/main.html"))
//var Tmls = template.Must(template.ParseGlob("../public/*"))


var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getStaticFiles(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()
	//Tmls.ExecuteTemplate(w, "main", nil)
	fmt.Fprintf(w, "The current directory: %s", dir)
}
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func mPageHandler(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "main.html", nil)

	if err != nil {
		return
	} else {
		fmt.Println("Sucessfully ")
	}
}

func addUser(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost {
		email := r.PostFormValue("email")
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		aUser := User{username, email, password }
		fmt.Fprintf(w, "Email: %s \n Username: %s \n Password: %s \n", aUser.Email, aUser.UserName, aUser.Password)
	}
	return
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/view/"):]
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	//title := r.URL.Path[len("/edit/"):]
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)

}

func saveHandler(w http.ResponseWriter, r *http.Request){
	//title := r.URL.Path[len("/save/"):]
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid page title")
	}
	return m[2], nil // The title is the second subexpression.
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}


func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	// Setup database
	db, err := sql.Open("mysql", "szhang:password@unix(/tmp/mysql.sock)/user?loc=Local")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	setupJSON()
	//setupCJWT()
	//a := AuthHandler{ db }
	//l := LoginHandler{ db}

	// this is tutorial from the golang.org
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	// main page for login and register
	http.HandleFunc("/main/", mPageHandler)

	// show how to serve static files
	http.HandleFunc("/getStaticFile", getStaticFiles)

	// pattern difference between register and login
	http.Handle("/register", registerHandler(db))
	http.Handle("/login", &LoginHandler{db })

    // demonstrate how to serve json
	http.HandleFunc("/getDrawing/", setupUsers )


	// Middleware to handle secure route
	http.HandleFunc("/verifyJWT", verifyJWT)
	http.Handle("/sendProtected", sendProtected())
	http.Handle("/adapt", Adapt(sendProtected(), isAuthenticated()))


	http.ListenAndServe(":8080", nil)

}
