package main

import (
	"fmt"
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
	http.HandleFunc("/", root)
	http.HandleFunc("/display", display)
	err := http.ListenAndServe(":"+getPort(), nil)
	if err != nil {
		panic(err)
	}
	// ListenAndServe never returns
	fmt.Println("listening...")

}

// will return even if no match is made
func root(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("default.html", "header.html", "footer.html")
	p := Data{Person{"Boris", ""}, time.Now()}
	t.Execute(w, p)
}

/*
func insertValues() {
}
*/

func display(w http.ResponseWriter, r *http.Request) {
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
