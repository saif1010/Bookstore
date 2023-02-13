package config
import (
	"github.com/jinzu/gorm"
	_ "github.com/jinzu/gorm/dialects/mysql"
)
var(
	db *gorm.DB
)
func Connects(){
	d,err:= gorm.Open("mysql","saif: saif@123/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err!=nill{
		panic(err)
	}
	db=d
}
func GetDB() *gorm.DB{
	return db
}