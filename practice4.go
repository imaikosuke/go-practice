package main

import (
    "fmt"
		"time"
    "net/http"
		"strings"
		"context"
    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// トークンを生成する関数
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
			Username: username,
			StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
			},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
			return "", err
	}

	return tokenString, nil
}

// トークンを検証する関数
func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
	})

	if err != nil {
			return nil, err
	}

	if !token.Valid {
			return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// ユーザー認証とJWTトークン発行のエンドポイント
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// 本来はユーザー名とパスワードを検証する必要がありますが、ここでは省略
	username := "testUser"
	tokenString, err := GenerateToken(username)
	if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
	}
	w.Write([]byte(tokenString))
}

// JWTトークンを検証するミドルウェア
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			// トークンはAuthorizationヘッダーから取得します
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
					http.Error(w, "Authorization header required", http.StatusUnauthorized)
					return
			}

			// Bearerスキーマを削除
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// トークンを検証
			claims, err := VerifyToken(tokenString)
			if err != nil {
					http.Error(w, "Invalid token", http.StatusUnauthorized)
					return
			}

			// トークンが有効な場合、リクエストコンテキストにユーザー名を追加
			ctx := context.WithValue(r.Context(), "username", claims.Username)
			next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// 認証済みユーザーのみアクセス許可するエンドポイント
func secretHandler(w http.ResponseWriter, r *http.Request) {
	// コンテキストからユーザー名を取得
	username := r.Context().Value("username")
	if username == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
	}
	fmt.Fprintf(w, "Welcome, %s", username)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/secret", authMiddleware(secretHandler))

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}