package main 

import (
	"fmt"
	"encoding/json"
	"net/http"
	//"strconv"
)

type Todo struct {
	ID 		int 	`json:"id"`
	Title 	string	`json:"title"`
	IsDone 	bool	`json:"is_done"`
}

func main() {
	todo := []Todo{
		{ID: 1,Title: "Goを復習する",IsDone: false},
	}

	http.HandleFunc("/todo",func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		if r.Method == http.MethodGet{
			json.NewEncoder(w).Encode(todo)
			return

		}
		if r.Method == http.MethodPost{
			var newTodo Todo
			err := json.NewDecoder(r.Body).Decode(&newTodo)
			if err != nil{
				http.Error(w, "Invaild JSON",http.StatusBadRequest)
				return
			}
		
		newTodo.ID	=len(todo) + 1

		todo = append(todo,newTodo)

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(newTodo)

		}

	})

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080",nil	)
	
}