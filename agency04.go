package main
import (
	
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"math/rand"
)

type tableElement struct {
	ID int `json:ID`
}

var tableElements1 [4][4]tableElement
var tableElements2 [4][4]tableElement

var ship = tableElement{ID:1}



func createTable(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&tableElements1)
	_ = json.NewDecoder(r.Body).Decode(&tableElements2)
	json.NewEncoder(w).Encode(tableElements1)
	json.NewEncoder(w).Encode(tableElements2)
}

func getTable(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w,"Player 1")
	json.NewEncoder(w).Encode(tableElements1)
	fmt.Fprintf(w, "Player 2")
	json.NewEncoder(w).Encode(tableElements2)
}


func updateTable(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	randomIndex1 := rand.Intn(4)
	randomIndex2 := rand.Intn(4)
	
  
	randomIndex3 := rand.Intn(4)
	randomIndex4 := rand.Intn(4)
   

	tableElements1[randomIndex1][randomIndex2] = tableElement{ID:1}
	tableElements2[randomIndex3][randomIndex4] = tableElement{ID:1}
	
	
	
	
}

func makeShot(w http.ResponseWriter, r *http.Request){
	
	for _, array1 := range tableElements1 {
		for i := range array1{
			//fmt.Print(array1[i], " ")
			if array1[i] == ship {
				fmt.Println("shot made from player 1")
				
			}
		} 
	}
	
	for _, array2 := range tableElements2 {
		for j := range array2{
			//fmt.Print(array2[j], " ")
			if array2[j] == ship {
				fmt.Println("shot made from player 2")
					
				

			}
		} 	
	}
	fmt.Println("........end of game turn")


} 

func main(){
	r := mux.NewRouter()

	


	r.HandleFunc("/player", getTable).Methods("GET")  
	r.HandleFunc("/player", createTable).Methods("POST")
	r.HandleFunc("/player", updateTable).Methods("PUT")
	r.HandleFunc("/player/shot", makeShot).Methods("PUT")


	log.Fatal(http.ListenAndServe(":8000", r))
}