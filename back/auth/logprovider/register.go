package logprovider

import (
	"back/auth/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)



func Register(c *gin.Context) {
    var user db.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request payload"})
        return
    }
	var found db.User
    err:=db.DB.Where("username = ?", user.Username).First(&found).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
        return
	}
	err=bcrypt.CompareHashAndPassword([]byte(found.Password),[]byte(user.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
        return
	}
	

}
