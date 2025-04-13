// @title TODO API
// @version 1.0
// @description シンプルなTODO管理API
// @host go-todo-hh8n.onrender.com
// @BasePath /
// @schemes https
package main

import (
	"go-todo/db"
	"go-todo/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "go-todo/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary TODO一覧取得
// @Description すべてのTODOを取得します
// @Produce json
// @Success 200 {array} models.Todo
// @Router /todos [get]
func main() {
    db.Init()
    r := gin.Default()

    // Swagger UI用のルート
    url := ginSwagger.URL("https://go-todo-hh8n.onrender.com/swagger/doc.json")
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    r.Static("/static", "./public")
    r.StaticFile("/", "./public/index.html")

    // 一覧取得
    r.GET("/todos", func(c *gin.Context) {
        var todos []models.Todo
        db.DB.Find(&todos)
        c.JSON(http.StatusOK, todos)
    })

    // 登録
    // @Summary TODOを登録
    // @Description 新しいTODOを作成します
    // @Accept json
    // @Produce json
    // @Param todo body models.Todo true "TODO情報"
    // @Success 200 {object} models.Todo
    // @Router /todos [post]
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
        port = "8080"
    }

    r.Run(":" + port)
}
