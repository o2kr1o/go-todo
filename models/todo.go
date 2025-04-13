package models

import "gorm.io/gorm"

// Todo Todoアイテムの情報
type Todo struct {
    ID        uint           `json:"id" example:"1"`
    Title     string         `json:"title" example:"買い物に行く"`
    Completed bool           `json:"completed" example:"false"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
}
