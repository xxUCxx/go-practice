package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	Age  int    `json:"age"`
}

func main() {
	users := []User{
		{ID: 1, Name: "Taro Cyber", Role: "Backend Engineer", Age: 21},
		{ID: 2, Name: "Hanako Ameba", Role: "Frontend Engineer", Age: 22},
	}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// ① GETリクエストなら一覧を返す
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(users)
			return
		}

		// ② POSTリクエストなら新しいユーザーを登録する
		if r.Method == http.MethodPost {
			var newUser User

			// ★空欄1：届いたJSON（r.Body）を newUser に読み込む（ステップ2の武器2を参照）
			err :=json.NewDecoder(r.Body).Decode(&newUser) 
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			// IDを自動で連番にする
			newUser.ID = len(users) + 1

			// ★空欄2：users スライスに newUser を追加する（ステップ2の武器3を参照）
			users =append(users,newUser) 

			// ★空欄3：登録した newUser をJSONで画面に返す（カンニングペーパーの鉄板パターン参照）
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newUser)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
