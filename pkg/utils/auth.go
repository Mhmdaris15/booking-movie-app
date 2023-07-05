package utils

import (
	"net/http"
	"time"

	"github.com/Mhmdaris15/booking-movie-app/internal/configs"
	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/pkg/database/mongo"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get The Email and Password

	var user models.User

	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Body"})
		return
	}

	// Hash The Password
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	db := mongo.ConnectDB()

	collection := mongo.GetCollection(db, "users")

	userFound := collection.FindOne(c, gin.H{
		"email": user.Email,
	})

	if userFound != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email Already Exist"})
		return
	}

	result, err := collection.InsertOne(c, gin.H{
		"username": user.Username,
		"name": user.Name,
		"email":    user.Email,
		"password": string(bcryptPassword),
		"age": user.Age,
		"balance": 0,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": result})
}

func Login(c *gin.Context) {
	// Get The Email and Password

	var body struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Body"})
		return
	}

	db := mongo.ConnectDB()

	var result models.User

	err := mongo.GetCollection(db, "users").FindOne(c, gin.H{
		"email": body.Email,
	}).Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": result})
}

func GenerateJWT() (string, error){
	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 30)
	claims["authorized"] = true
	claims["user"] = "username"

	tokenString, err := token.SignedString([]byte(configs.SecretKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyJWT(endpointHandler func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.SecretKey()), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user", claims["user"])
			endpointHandler(c)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
	}
}