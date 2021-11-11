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

var noship = tableElement{ID:0}

type Position struct {
	El1 int `json:el1`
	El2 int	`json:el2`
}

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
/* 
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


}  */

func makeShotPlayer1(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var positionPlayer1 Position
	_ = json.NewDecoder(r.Body).Decode(&positionPlayer1)
	positionX := positionPlayer1.El1
	positionY := positionPlayer1.El2
	if (tableElements1[positionX][positionY] == ship) {
		tableElements2[positionX][positionY] = noship
	}
	fmt.Println("shot made from player1")
	fmt.Println(positionX)
	fmt.Println(positionY)

}

func makeShotPlayer2(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var positionPlayer2 Position
	_ = json.NewDecoder(r.Body).Decode(&positionPlayer2)
	positionX := positionPlayer2.El1
	positionY := positionPlayer2.El2
	if (tableElements2[positionX][positionY] == ship) {
		tableElements1[positionX][positionY] = noship
	}
	fmt.Println("shot made from player2")
	fmt.Println(positionX)
	fmt.Println(positionY)

}


func main(){
	r := mux.NewRouter()
	
		


	r.HandleFunc("/player", getTable).Methods("GET")  
	r.HandleFunc("/player", createTable).Methods("POST")
	r.HandleFunc("/player", updateTable).Methods("PUT")
	//r.HandleFunc("/player/shot", makeShot).Methods("PUT")
	r.HandleFunc("/player1/shot", makeShotPlayer1).Methods("POST")
	r.HandleFunc("/player2/shot", makeShotPlayer2).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}