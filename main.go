package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	Age  int    `json:"age"`
}

func main() {
	// 1. 複数のユーザーデータ（スライス）を用意
	users := []User{
		{ID: 1, Name: "Taro Cyber", Role: "Backend Engineer", Age: 21},
		{ID: 2, Name: "Hanako Ameba", Role: "Frontend Engineer", Age: 22},
		{ID: 3, Name: "Ken Abema", Role: "Infrastructure Engineer", Age: 23},
	}

	// 窓口①：全ユーザー一覧を返す (/users)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	// 窓口②：指定したIDのユーザー1人を返す (/user?id=2)
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// URLの「?id=...」の文字を取り出し、数字（int）に変換する
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// スライスの中からIDが一致するユーザーを1人ずつ探す
		for _, u := range users {
			if u.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(u)
				return
			}
		}

		// 最後まで見つからなかった場合は 404 エラーを返す
		http.Error(w, "User not found", http.StatusNotFound)
	})

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
