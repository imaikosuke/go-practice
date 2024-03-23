package main

import (
    "fmt"
    "time"
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

func main() {
    // トークン生成のテスト
    tokenString, err := GenerateToken("username")
    if err != nil {
        fmt.Println("Token generation error:", err)
        return
    }
    fmt.Println("Generated Token:", tokenString)

    // トークン検証のテスト
    claims, err := VerifyToken(tokenString)
    if err != nil {
        fmt.Println("Token verification error:", err)
        return
    }
    fmt.Println("Token verified for user:", claims.Username)
}
