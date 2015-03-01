package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
)

type Task struct {
	Name      string
	Frequency string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("Tasks")
	/*
		err = c.Insert(&Task{"Thumbs", "daily"},
			&Task{"Feet", "weekly"})
		if err != nil {
			log.Fatal(err)
		}
	*/

	result := []Task{}
	//err = c.Find(bson.M{"name": "Thumbs"}).One(&result)
	err = c.Find(nil).All(&result)
	fmt.Println("Task:", result)
	if err != nil {
		log.Fatal(err)
	}

	for x, y := range result {
		fmt.Println("Task:", x, y.Name)
	}
}
