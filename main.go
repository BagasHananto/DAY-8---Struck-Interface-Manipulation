package main

import (
	"fmt"

	"log"

	"strconv"

	"net/http"

	"github.com/gorilla/mux"

	"html/template"
)

var Data = map[string]interface{}{
	"Title": "Personal Web",
}

type Project struct {
	Title       string
	Post_date   string
	Start_date  string
	End_date    string
	Duration    string
	Author      string
	Description string
	NodeJs      string
	Java        string
	Php         string
	Laravel     string
}

var Projects = []Project{
	{
		Title:       "Pembelajaran Online",
		Duration:    "Duration : 3 Weeks",
		Author:      " | Bagas",
		Description: "Sangat sulit sekali hehehehe",
	},
}

// function routing
func main() {
	router := mux.NewRouter()

	// Create Folder
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/Project", project).Methods("GET")
	router.HandleFunc("/addProject", addProject).Methods("POST")
	router.HandleFunc("/contactMe", contactMe).Methods("GET")
	router.HandleFunc("/projectDetail/{id}", projectDetail).Methods("GET")
	router.HandleFunc("/delete-project/{id}", deleteProject).Methods("GET")

	fmt.Println("Server Running Successfully")
	http.ListenAndServe("localhost:5000", router)
}

// function handling index.html
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Parsing template html file
	var tmpl, err = template.ParseFiles("index.html")
	//Error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	resp := map[string]interface{}{
		"Title":    Data,
		"Projects": Projects,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}

// function handling myproject.html
func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Parsing template html file
	var tmpl, err = template.ParseFiles("myproject.html")
	//Error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

// function handling contactMe.html
func contactMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Parsing template html file
	var tmpl, err = template.ParseFiles("contactMe.html")
	//Error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

// function handling myproiect-detail.html
func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//Parsing template html file
	var tmpl, err = template.ParseFiles("myproject-detail.html")
	//Error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	ProjectDetail := Project{}

	for i, data := range Projects {
		if i == id {
			ProjectDetail = Project{
				Title:       data.Title,
				Start_date:  data.Start_date,
				End_date:    data.End_date,
				Post_date:   data.Post_date,
				Description: data.Description,
			}
		}
	}

	resp := map[string]interface{}{
		"Data":    Data,
		"Project": ProjectDetail,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	title := r.PostForm.Get("title")
	description := r.PostForm.Get("description")
	startDate := r.PostForm.Get("start-date")
	endDate := r.PostForm.Get("end-date")
	nodeJs := r.PostForm.Get("NodeJs")
	java := r.PostForm.Get("Java")
	php := r.PostForm.Get("Php")
	laravel := r.PostForm.Get("Laravel")

	var newProject = Project{
		Title:       title,
		Start_date:  startDate,
		End_date:    endDate,
		Description: description,
		NodeJs:      nodeJs,
		Java:        java,
		Php:         php,
		Laravel:     laravel,
	}

	Projects = append(Projects, newProject)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	Projects = append(Projects[:id], Projects[id+1:]...)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
