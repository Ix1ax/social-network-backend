package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ix1ax/social-network-backend/internal/user/dto"
	"github.com/ix1ax/social-network-backend/internal/user/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userSevice service.UserService) *UserHandler {
	return &UserHandler{userService: userSevice}
}

func (h *UserHandler) RegisterRoutes(rg *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	auth := rg.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}
	users := rg.Group("/users")
	users.Use(authMiddleware)
	{
		users.GET("/me", h.GetMe)
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.userService.Register(c.Request.Context(), &req)

	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)

}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.userService.Login(c.Request.Context(), &req)

	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}

func (h *UserHandler) GetMe(c *gin.Context) {

	userIDVal, exists := c.Get("userID")

	if exists == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	res, err := h.userService.GetByID(c.Request.Context(), userIDVal.(uuid.UUID))

	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}
