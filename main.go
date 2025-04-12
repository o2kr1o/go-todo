package main

import (
    "go-todo/db"
    "go-todo/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
)

func main() {
    db.Init()
    r := gin.Default()
    r.Static("/static", "./public")
    r.StaticFile("/", "./public/index.html")

    // 一覧取得
    r.GET("/todos", func(c *gin.Context) {
        var todos []models.Todo
        db.DB.Find(&todos)
        c.JSON(http.StatusOK, todos)
    })

    // 登録
    r.POST("/todos", func(c *gin.Context) {
        var todo models.Todo
        if err := c.ShouldBindJSON(&todo); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        db.DB.Create(&todo)
        c.JSON(http.StatusOK, todo)
    })

    // 更新
    r.PUT("/todos/:id", func(c *gin.Context) {
        var todo models.Todo
        id := c.Param("id")
        if err := db.DB.First(&todo, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
            return
        }
        c.ShouldBindJSON(&todo)
        db.DB.Save(&todo)
        c.JSON(http.StatusOK, todo)
    })

    // 削除
    r.DELETE("/todos/:id", func(c *gin.Context) {
        id := c.Param("id")
        db.DB.Delete(&models.Todo{}, id)
        c.Status(http.StatusNoContent)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // ローカル開発用デフォルト
    }

    r.Run(":" + port)
}
