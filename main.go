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
    // @Summary TODO一覧取得
    // @Description すべてのTODOを取得します
    // @Produce json
    // @Success 200 {array} models.Todo
    // @Router /todos [get]
    r.GET("/todos", getTodos)

    // 登録
    // @Summary TODOを登録
    // @Description 新しいTODOを作成します
    // @Accept json
    // @Produce json
    // @Param todo body models.Todo true "TODO情報"
    // @Success 201 {object} models.Todo
    // @Failure 400 {object} map[string]string "無効なリクエスト"
    // @Router /todos [post]
    r.POST("/todos", createTodo)

    // 更新
    // @Summary TODOを更新
    // @Description 指定したIDのTODOを更新します
    // @Accept json
    // @Produce json
    // @Param id path int true "TODO ID"
    // @Param todo body models.Todo true "更新内容"
    // @Success 200 {object} models.Todo
    // @Failure 400 {object} map[string]string "無効なリクエスト"
    // @Failure 404 {object} map[string]string "TODOが見つかりません"
    // @Router /todos/{id} [put]
    r.PUT("/todos/:id", updateTodo)

    // 削除
    // @Summary TODOを削除
    // @Description 指定したIDのTODOを削除します
    // @Param id path int true "TODO ID"
    // @Success 204 "削除成功"
    // @Failure 404 {object} map[string]string "TODOが見つかりません"
    // @Router /todos/{id} [delete]
    r.DELETE("/todos/:id", deleteTodo)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(":" + port)
}

func getTodos(c *gin.Context) {
    var todos []models.Todo
    db.DB.Find(&todos)
    c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Create(&todo)
    c.JSON(http.StatusCreated, todo)
}

func updateTodo(c *gin.Context) {
    var todo models.Todo
    id := c.Param("id")
    
    if err := db.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
        return
    }
    
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    db.DB.Save(&todo)
    c.JSON(http.StatusOK, todo)
}

func deleteTodo(c *gin.Context) {
    id := c.Param("id")
    result := db.DB.Delete(&models.Todo{}, id)
    
    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
        return
    }
    
    c.Status(http.StatusNoContent)
}
