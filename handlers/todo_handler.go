package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"todo/models"
)
func CreateTodo(w http.ResponseWriter,r *http.Request){ 
	var todo models.Todo
	err:= json.NewDecoder(r.Body).Decode(&todo)
	if err != nil{
		http.Error(w,"Invalid req", http.StatusBadRequest)
		return

	}
	todo.Completed=false
	err=mgm.Coll(&todo).Create(&todo)
	if err!=nil{
		http.Error(w,"Failed to create todo", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
func GetTodos(w http.ResponseWriter , r *http.Request){
	var todos []models.Todo
	err := mgm.Coll(&models.Todo{}).SimpleFind(&todos,bson.M{})
	if err != nil{
		http.Error(w,"Failed to fetch todos" , http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}
func GetTodoByID(w http.ResponseWriter , r *http.Request){
	params := mux.Vars(r)
	id,err := primitive.ObjectIDFromHex(params["id"])
	if err!=nil{
		http.Error(w,"ID not found", http.StatusBadRequest)
	}
	todo := &models.Todo{}
	err = mgm.Coll(todo).FindByID(id ,todo)
	if err!=nil{
		http.Error(w,"Todo not found" , http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(todo)
}
//// Update a Todo
func UpdateTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updateData map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	todo := &models.Todo{}
	err = mgm.Coll(todo).FindByID(id, todo)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Update fields if provided
	if title, ok := updateData["title"].(string); ok {
		todo.Title = title
	}
	if completed, ok := updateData["completed"].(bool); ok {
		todo.Completed = completed
	}

	err = mgm.Coll(todo).Update(todo)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

// Delete a Todo
func DeleteTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	todo := &models.Todo{}
	err = mgm.Coll(todo).FindByID(id, todo)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	err = mgm.Coll(todo).Delete(todo)
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
