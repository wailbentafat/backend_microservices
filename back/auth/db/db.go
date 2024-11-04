package db

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
type User struct{
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}
func Migrate(db *gorm.DB){
	db.AutoMigrate(&User{},&Key{})
}
func Init(db *gorm.DB)(*gorm.DB){
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
    if err != nil {
        fmt.Println("Error opening database:", err)
        return nil
    }
 DB=db
 Migrate(db)
 return db
}
type Key struct{
	Keyval string `gorm:"unique`
    Active bool `gorm:"default:true"`
    User_id uint`gorm:"index"`
	User User `gorm:"foreignKey:User_id"`
	Created_at time.Time

}