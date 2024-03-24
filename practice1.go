package main  // パッケージ名をmainにすることで、実行可能なプログラムを作成することができる

import (
	"fmt"       // 標準入出力に使用
	"net/http"  // HTTPサーバーを作成するために使用
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // ルートパスにアクセスがあった場合の処理
		fmt.Fprintf(w, "Hello HTTP World!")                 // レスポンスを返す
	})

	fmt.Println("http://localhost:8080 でサーバーを起動します。")   // サーバーを起動する旨を表示
	if err := http.ListenAndServe(":8080", nil); err != nil {  	// サーバーを起動
		panic(err)                                                // エラーが発生した場合はエラー内容を表示
	}
}