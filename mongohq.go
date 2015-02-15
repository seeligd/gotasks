package main

import (
	"fmt"
	"html/template"
	"net/http"
	//"os"
	"time"
)

type Data struct {
	Name  string
	Email string
	Time  time.Time
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/d", display)
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
	p := Data{"Boris", "boris@yeltsin.gov", time.Now()}
	t.Execute(w, p)
}

func display(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hellow")
}

func getPort() string {
	port := "3001"
	fmt.Println("Now Serving on port", port)
	return port
}
