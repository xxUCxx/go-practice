package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 1. 構造体（Struct）：ユーザーデータの設計図
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
	Age      int    `json:"age"`
}

func main() {
	// "/user" というURLにアクセスが来たときの処理
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// 2. 設計図（User）をもとに実際のデータを作成（:= で変数定義）
		user := User{
			ID:       1,
			Name:     "maki",
			Role:     "Backend Engineer",
			IsActive: true,
			Age:      20,
		}

		// 3. レスポンスをJSON形式に設定して送信
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
