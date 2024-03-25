package main

import (
    "log"
    "net/http"
    "time"
)

// ミドルウェア関数
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        // リクエストを次のハンドラに渡す
        next.ServeHTTP(w, r)
        // 処理時間を計測してログに記録
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

func main() {
    // 通常のハンドラ
    myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))
    })

    // ミドルウェアを組み込んだハンドラを作成
    wrappedHandler := LoggingMiddleware(myHandler)

    http.Handle("/", wrappedHandler)
    http.ListenAndServe(":8080", nil)
}
