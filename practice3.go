package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

type User struct {
    gorm.Model
    Name  string
    Email string `gorm:"type:varchar(100);unique_index"`
}

func main() {
    dsn := "host=localhost user=imaikosuke dbname=go-practice-database sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("データベース接続に失敗しました: %v", err)
    }

    // モデルに基づいてテーブルを作成
    db.AutoMigrate(&User{})

    // ユーザーを作成
    db.Create(&User{Name: "Taro", Email: "taro@example.com"})

    // ユーザーを取得
    var user User
    db.First(&user, "name = ?", "Taro")
    log.Printf("取得したユーザー: %#v", user)
}
