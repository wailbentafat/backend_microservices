package logprovider

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"back/auth/db"
)


func login(c *gin.Context) {
    var user db.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request payload"})
        return
    }
    fmt.Println(user)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return
	}
	user.Password = string(hash)
	db.DB.Create(&user)
	ID := user.ID
	ha, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%d", ID)), 5)
	if err != nil {
		return
	}
	keyval:=fmt.Sprintf("%xwi%Eth$", ha,user.Username)
	key := db.Key{
		Keyval: keyval,
		User_id: ID,
		Active: true,
		Created_at: time.Now(),
	}
	db.DB.Create(&key)

	c.JSON(200, gin.H{"message": "User created successfully"})
}