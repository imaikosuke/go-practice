package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "sync"
		"strconv"
    "github.com/gorilla/mux"
)

// メモリ上に保存するデータ構造
type Memo struct {
    ID   int    `json:"id"`
    Text string `json:"text"`
}

// メモリ上に保存するデータ
var (
    memos []Memo
    mu    sync.Mutex
    idSeq int
)

// メモリ上に保存するデータを初期化
func getMemosHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(memos)
}

// メモリ上に保存するデータを追加
func createMemoHandler(w http.ResponseWriter, r *http.Request) {
    var memo Memo
    if err := json.NewDecoder(r.Body).Decode(&memo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    mu.Lock()
    idSeq++
    memo.ID = idSeq
    memos = append(memos, memo)
    mu.Unlock()

    w.Header().Set("Location", fmt.Sprintf("/memos/%d", memo.ID))
    w.WriteHeader(http.StatusCreated)
}

// メモを削除するハンドラー関数
func deleteMemoHandler(w http.ResponseWriter, r *http.Request) {
	// URLからメモのIDを取得
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, memo := range memos {
			if memo.ID == id {
					// メモを削除
					memos = append(memos[:i], memos[i+1:]...)
					w.WriteHeader(http.StatusNoContent) // 204 No Content
					return
			}
	}

	// メモが見つからなかった場合
	http.Error(w, "Memo not found", http.StatusNotFound)
}


func main() {
    r := mux.NewRouter()
    r.HandleFunc("/memos", getMemosHandler).Methods("GET")
    r.HandleFunc("/memos", createMemoHandler).Methods("POST")
		r.HandleFunc("/memos/{id:[0-9]+}", deleteMemoHandler).Methods("DELETE")

    fmt.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
