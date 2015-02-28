package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
	"time"
)

type Data struct {
	Person Person
	Time   time.Time
}
type Person struct {
	Name  string
	Email string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainPage)
	/*
		r.HandleFunc("/tasks/{taskId}", viewTask)
		r.HandleFunc("/task/new", createTask)
		r.HandleFunc("/task/edit", editTask)
	*/

	http.Handle("/", r) // give everything to gorilla
	err := http.ListenAndServe(":"+getPort(), nil)
	if err != nil {
		panic(err)
	}
	// ListenAndServe never returns
	fmt.Println("listening...")

}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("default.html", "header.html", "footer.html")
	p := Data{Person{"Boris", ""}, time.Now()}
	t.Execute(w, p)
}

func viewTask(w http.ResponseWriter, r *http.Request) {
	displayTemplate, _ := template.ParseFiles("display.html", "header.html", "footer.html")
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	collection := session.DB("test").C("People")
	result := Person{}
	err = collection.Find(bson.M{"name": r.FormValue("name")}).One(&result)

	fmt.Println(err)

	data := Data{result, time.Now()}
	if result.Email != "" {
		errn := displayTemplate.Execute(w, data)
		if errn != nil {
			http.Error(w, errn.Error(), http.StatusInternalServerError)
		}
	} else {
		displayTemplate.Execute(w, data)
	}
}

func getPort() string {
	port := "3001"
	fmt.Println("Now Serving on port", port)
	return port
}
