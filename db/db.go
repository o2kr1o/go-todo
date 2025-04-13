package db

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "go-todo/models"
    "os"
)

var DB *gorm.DB

func Init() {
    dbPath := "todo.db"
    if os.Getenv("GIN_MODE") == "release" {
        dbPath = "/tmp/todo.db"
    }
    
    var err error
    DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        log.Fatal("DB接続失敗:", err)
    }
    DB.AutoMigrate(&models.Todo{})
    
    var count int64
    DB.Model(&models.Todo{}).Count(&count)
    if count == 0 {
        DB.Create(&models.Todo{Title: "サンプルTODO", Completed: false})
    }
}
