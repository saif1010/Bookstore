package models
import (
	"github.com/jinzu/gorm"
	"github.com/saif1010/go-bookstore/pkg/config"
)
var db *gorm.DB
type Book struct{
	gorm.model 
	Name string`gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}
func init(){
	config.Connects()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}