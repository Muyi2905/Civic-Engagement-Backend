package controllers

import (
	"muyi2905/civic/backend/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var validate = validator.New()
var jwtSecret = []byte(os.Getenv("JWT-SECRET"))

type Claims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}


func Signup(db *gorm.DB, c *gin.Context) {
	
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}


	if err := validate.Struct(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password hashing failed"})
		return
	}
	user.Password = string(hashedPassword)

	
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	// Generate JWT token
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}


func Login(db *gorm.DB, c *gin.Context){
	
	var user models.User
	if err:= db.First(&user).Find("id")
}

func GetUsers(db *gorm.DB, c *gin.Context) {
	var user models.User
	if err := db.Find(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to get users"})
	}
	c.JSON(http.StatusOK, gin.H{"users": user})
}

func GetUserById(db *gorm.DB, c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "userid/user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	}
}




