package main
import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	
)
type Person struct{
	Id  string   `json:"id,omitempty"`
	Name string `json:"nome"`
	Age string `json:"idade"`
    Adress string `json:"endereço`

}
var people []Person

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", GetPerson).Methods("GET")
	r.HandleFunc("/person/{id}", GetPersonId).Methods("GET")
    r.HandleFunc("/person/{id}", UpdatePerson).Methods("PUT")
    r.HandleFunc("/person", CreatePerson).Methods("POST")
	r.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")

    http.ListenAndServe(":8080", r)
}


func GetPersonId(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range people {
        if item.Id == params["id"] && item.Id != ""{
            fmt.Println(item)
            json.NewEncoder(w).Encode(item)
        }
    }    
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(people)
}
 
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    people = append(people, person)
    json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range people {
        if item.Id == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")          
    params := mux.Vars(r)
    for index, item := range people {
        if item.Id == params["id"] {
            people = append(people[:index], people[index+1:]...)
            var person Person
            _ = json.NewDecoder(r.Body).Decode(&person)
            people = append(people, person)
            json.NewEncoder(w).Encode(&person)
        
            return
        }
    }
    json.NewEncoder(w).Encode(people)

}
