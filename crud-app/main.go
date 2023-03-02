package main

import (
	"fmt"
	//	"log"
	// "encoding/json"
	//"math/rand"
	//"net/http"
	//"strconv"
	//"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func main() {
	fmt.Println("CRUD WITH MOVIES")
}
