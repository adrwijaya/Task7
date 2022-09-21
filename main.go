package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	route.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	route.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	route.PathPrefix("/icon/").Handler(http.StripPrefix("/icon/", http.FileServer(http.Dir("./icon"))))

	route.HandleFunc("/home", home).Methods("GET")
	route.HandleFunc("/add_myproject", addMyProject).Methods("GET")
	route.HandleFunc("/contact_me", contactMe).Methods("GET")
	route.HandleFunc("/detail_project", detailProject).Methods("GET")
	route.HandleFunc("/add_myproject", ambilData).Methods("POST")

	fmt.Println("server running on port 5000")
	http.ListenAndServe("localhost:80", route)

}

func ambilData(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Project Name : " + r.PostForm.Get("project-name"))
	fmt.Println("Start Date : " + r.PostForm.Get("start-date"))
	fmt.Println("End Date : " + r.PostForm.Get("end-date"))
	fmt.Println("Description : " + r.PostForm.Get("description"))
	fmt.Println("Technologies : " + r.PostForm.Get("tech"))
	fmt.Println("Technologies : " + r.PostForm.Get("tech2"))
	fmt.Println("Technologies : " + r.PostForm.Get("tech3"))
	fmt.Println("Technologies : " + r.PostForm.Get("tech4"))

	http.Redirect(w, r, "/add_myproject", http.StatusMovedPermanently)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("html/index.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func addMyProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("html/add_myproject.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func contactMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("html/contact_me.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func detailProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("html/detail_project.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}
