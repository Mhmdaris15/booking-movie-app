package handlers

import (
	"net/http"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetAllUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	GetUserByEmail(ctx *gin.Context)
	GetUserByUsername(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	UpdateUserByUsername(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (h *userHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.userService.GetUserByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := h.userService.GetUserByEmail(ctx, email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := h.userService.GetUserByUsername(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) CreateUser(ctx *gin.Context) {
	var user *models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.userService.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	

	ctx.JSON(http.StatusCreated, user)
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	var user *models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.userService.UpdateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) UpdateUserByUsername(ctx *gin.Context) {
	// username := ctx.Param("username")
	var user *models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.userService.UpdateUserByUsername(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.userService.DeleteUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

// Path: internal\handlers\showtime_handler.go