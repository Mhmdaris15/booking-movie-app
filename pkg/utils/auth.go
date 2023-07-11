package utils

import (
	"net/http"
	"time"

	"github.com/Mhmdaris15/booking-movie-app/internal/configs"
	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/pkg/database/mongodb"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenExpiration = 2 * time.Hour
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	// Role string `json:"role"`
	jwt.StandardClaims
}

func Signup(c *gin.Context) {
	// Get the Email and Password

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Body"})
		return
	}

	// Hash the Password
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	db := mongodb.ConnectDB()

	collection := mongodb.GetCollection(db, "users")

	// Check if user with the same email already exists
	err = collection.FindOne(c, bson.M{"email": user.Email}).Err()
	if err != nil && err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	} else if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email Already Exists"})
		return
	}

	result, err := collection.InsertOne(c, bson.M{
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
		"password": string(bcryptPassword),
		"age":      user.Age,
		"balance":  0,
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

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Body"})
		return
	}

	db := mongodb.ConnectDB()

	var result models.User

	collection := mongodb.GetCollection(db, "users")
	if collection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	err := collection.FindOne(c, bson.M{"email": body.Email}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Email or Password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Email or Password"})
		return
	}

	token, err := GenerateJWT(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": result, "token": token})
}

func GenerateJWT(user models.User) (string, error) {
	claims := JWTClaims{
		user.ID.Hex(),
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpiration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(configs.SecretKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the authorization header
		authHeader := c.GetHeader("Authorization")

		// Check if the authorization header is present
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Header Required"})
			c.Abort()
			return
		}

		// Extract the token from the authorization header
		tokenString := authHeader[len("Bearer "):] // Remove the Bearer prefix

		// Parse the JWT token
		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.SecretKey()), nil
		})

		// Check if there was an error in parsing JWT
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": `Invalid Token | when parsing JWT`, "error": err.Error()})
			c.Abort()
			return
		}

		// Verify the token claims
		if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)
			// c.Set("role", claims.Role)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token | in claiming Token", "error": err.Error()})
			c.Abort()
			return
		}
	}
}

func ProtectedHandler(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")
	// role, _ := c.Get("role")

	c.JSON(http.StatusOK, gin.H{"message": "Protected Route", "user_id": userID, "email": email})

}