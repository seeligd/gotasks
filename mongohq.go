package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
)

type E interface{}

type Data struct {
	E     // anonymous field; 'Data has an E'
	Title string
}

type Task struct {
	Name      string
	Frequency string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", tasksPage)
	r.HandleFunc("/", mainPage)

	r.HandleFunc("/task/new", editTask)
	r.HandleFunc("/task", createTask).Methods("POST")

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
	t.Execute(w, Data{nil, "Task App"})
}

func getDb() (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("Tasks")

	return c, session
}

// show all tasks
func tasksPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tasks.html", "header.html", "footer.html")
	//p := []Task{Task{"cut toenails", "monthly"}, Task{"eat carrots", "weekly"}}

	collection, session := getDb()
	defer session.Close()

	result := []Task{}
	err := collection.Find(nil).All(&result)

	for _, x := range result {
		fmt.Println(x)
	}
	if err != nil {
		panic(err)
	}

	//result := Task{}
	//query(&bson.M{}, E(result))
	t.Execute(w, Data{result, "title"})
}

// show specific task
func viewTask(w http.ResponseWriter, r *http.Request) {
	displayTemplate, _ := template.ParseFiles("display.html", "header.html", "footer.html")

	collection, session := getDb()
	defer session.Close()
	result := Task{}
	err := collection.Find(bson.M{"name": r.FormValue("name")}).One(&result)

	fmt.Println(err)

	data := Data{result, "title"}
	displayTemplate.Execute(w, data)
	/**
	if result.Email != "" {
		errn := displayTemplate.Execute(w, data)
		if errn != nil {
			http.Error(w, errn.Error(), http.StatusInternalServerError)
		}
	} else {
		displayTemplate.Execute(w, data)
	}
	**/
}

func editTask(w http.ResponseWriter, r *http.Request) {
	displayTemplate, _ := template.ParseFiles("editTask.html", "header.html", "footer.html")
	displayTemplate.Execute(w, nil)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	displayTemplate, _ := template.ParseFiles("editTask.html", "header.html", "footer.html")

	collection, session := getDb()
	defer session.Close()
	task := Task{r.FormValue("name"), r.FormValue("frequency")}
	err := collection.Insert(&task)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
	displayTemplate.Execute(w, nil)
}

func getPort() string {
	port := "3001"
	fmt.Println("Now Serving on port", port)
	return port
}
