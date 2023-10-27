package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

var db *sql.DB
var store = sessions.NewCookieStore([]byte("super-secret-key"))

type User struct {
	Login  string
	Password string
	FirstName string
}

func main() {

	var err error
	db, err = sql.Open("postgres", "user=user password=111 dbname=webappdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/", mainPageHandler)
	mux.HandleFunc("/main", mainPageHandler)
	mux.HandleFunc("/login", loginHandler)

	mux.HandleFunc("/favicon.ico", faviconHandler)

	log.Println("Start server")
	err = http.ListenAndServe("localhost:8080", mux)
	log.Fatal(err)
}

func loginHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		ts, err := template.ParseFiles("resources/html/login.html")

		if err != nil {
			log.Println(err.Error())
			http.Error(writer, "Internal Server Error", 500)
			return
		}

		err = ts.Execute(writer, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(writer, "Internal Server Error", 500)
		}

		return
	}

	if request.Method == http.MethodPost {
		login := request.FormValue("loginInput")
		password := request.FormValue("passwordInput")

		login = strings.TrimSpace(login)
		password = strings.TrimSpace(password)

		if len(login) != 0 {
			rows := db.QueryRow("SELECT Login, Password, FirstName FROM UserEntity WHERE Login = $1", login)

			user := &User{}

			err := rows.Scan(&user.Login, &user.Password, &user.FirstName)
			if err != nil{
				log.Println(err.Error())
			}

			log.Println(user.Login, user.FirstName)

			if password == user.Password {

				session, err1 := store.Get(request, "session-name")

				if err1 != nil {
					http.Error(writer, err.Error(), http.StatusInternalServerError)
					return
				}

				// Set some session values.
				session.Values["authenticated"] = true
				// Save it before we write to the response/return from the handler.
				err2 := session.Save(request, writer)

				log.Println("success login!")

				if err2 != nil {
					http.Error(writer, err2.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(writer, request, "/main", http.StatusSeeOther)

			} else {
				log.Println("fail login!")
				http.Redirect(writer, request, "/login", http.StatusSeeOther)
			}

		}

	}
}

func mainPageHandler(writer http.ResponseWriter, request *http.Request) {

	session, err1 := store.Get(request, "session-name")
	if err1 != nil {
		http.Error(writer, err1.Error(), http.StatusInternalServerError)
		return
	}

	auth, ok := session.Values["authenticated"].(bool)

	if !ok || !auth {
		http.Error(writer, "Forbidden", http.StatusForbidden)
        return
	}

	_, err2 := writer.Write([]byte("Main page handler"))

	if err2 != nil {
		log.Print(err2)
	}
}

func faviconHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "resources/favicon.ico")
}