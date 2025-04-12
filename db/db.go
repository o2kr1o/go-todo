package db

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "go-todo/models"
)

var DB *gorm.DB

func Init() {
    var err error
    DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("DB接続失敗:", err)
    }
    DB.AutoMigrate(&models.Todo{})
}
