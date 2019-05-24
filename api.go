//main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Angka struct
type Angka struct {
	Number int `json:"number"`
}

//Text struct
type Text struct {
	Teks string `json:"text"`
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/spell", spellFunc).Methods("GET")
	router.HandleFunc("/read", readFunc)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page")
	fmt.Println("Endpoint Hit: HomePage")
}

func spellFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("spell")
	params, _ := r.URL.Query()["number"]
	// fmt.Println(params)
	str := params[0]
	// fmt.Println(str)
	num, _ := strconv.Atoi(str)
	res := spellAngka(num)
	fmt.Println(res)
	ret := Text{Teks: res}
	json.NewEncoder(w).Encode(ret)
	// respondWithJSON(w, http.statusOK)
}

func readFunc(w http.ResponseWriter, r *http.Request) {
	var baca Text
	fmt.Println("read")
	json.NewDecoder(r.Body).Decode(&baca) //catch data from post req and Post to our struct model
	fmt.Println(baca)
	res := prosesString(baca.Teks)

	fmt.Println(baca.Teks)
	fmt.Println(res)
	ret := Angka{Number: res}
	json.NewEncoder(w).Encode(ret)
}

func main() {
	handleRequests()
}
