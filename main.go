package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	//"strconv"
)
type  Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}
var todos []Todo
var index int = 1;

func createTodo(w http.ResponseWriter,r *http.Request){
	var todo Todo
	err:=json.NewDecoder(r.Body).Decode(&todo)
	if err!=nil{
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	todo.ID=index
	todo.Completed = false
	index++;
	todos = append(todos,todo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)


}
func getTodo(w http.ResponseWriter,r *http.Request){
	json.NewEncoder(w).Encode(todos)
}
func getTodobById (w http.ResponseWriter, r *http.Request){
	params:= mux.Vars(r)
	id,err := strconv.Atoi(params["id"])
	if err != nil{
		http.Error(w , "Invalid Request payload", http.StatusBadRequest)

	}
	for _,todo := range todos{
		if todo.ID==id{
			json.NewEncoder(w).Encode(todo)
			return
		}
	}	
	http.Error(w, "todo not found",http.StatusNotFound)

}

func updateById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"]) // Convert ID to integer
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// Decode only the fields that are provided in the request body
	var updatedFields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updatedFields)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Find and update the existing todo item
	for i, todo := range todos {
		if todo.ID == id {
			// Only update provided fields
			if title, ok := updatedFields["title"].(string); ok {
				todos[i].Title = title
			}
			if completed, ok := updatedFields["completed"].(bool); ok {
				todos[i].Completed = completed
			}

			// Respond with the updated todo
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}

	http.Error(w, "ID not found", http.StatusNotFound)
}

func deleteById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"]) // Convert the ID to integer
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// Find the index of the item to delete
	index := -1
	for i, todo := range todos {
		if todo.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Remove the item from the slice
	todos = append(todos[:index], todos[index+1:]...)

	w.WriteHeader(http.StatusNoContent) // 204 No Content (successful delete)
}


func main(){

	r := mux.NewRouter()
	r.HandleFunc("/todo",createTodo).Methods("POST")
	r.HandleFunc("/todo",getTodo).Methods("GET")
	r.HandleFunc("/updatetodo/{id}",updateById).Methods("PUT")
	r.HandleFunc("/todo/{id}",getTodobById).Methods("GET")
	r.HandleFunc("/todo/{id}",deleteById).Methods("DELETE")

	fmt.Println("Server runnning at 3001")
	http.ListenAndServe(":3001",r)

//	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		//fmt.Fprintf(w,"Hello")
	//})
	//http.HandleFunc("/hello",func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w,"Hello from hello path")
	//})
	//log.Fatal(http.ListenAndServe(":3001",nil))
	

}