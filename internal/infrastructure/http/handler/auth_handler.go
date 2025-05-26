package handler

import (
	"net/http"
	userApp "userman/internal/application/user"
	"userman/internal/infrastructure/jwt"
	authHttp "userman/internal/interface/http/auth"
	"userman/internal/interface/http/auth/dto"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	userService userApp.UserService
	jwtService  jwt.JWTService
}

func NewAuthHandler(userService userApp.UserService, jwtService jwt.JWTService) authHttp.AuthHandler {
	return &authHandler{userService: userService, jwtService: jwtService}
}

func (h *authHandler) Register(c *gin.Context) {
	var body dto.RegisterBodyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.CreateUser(c.Request.Context(), body.Name, body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, err := h.jwtService.GenerateToken(result.ID, result.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.MapTokenToDto(token))
}

func (h *authHandler) Login(c *gin.Context) {
	var body dto.LoginBodyRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.GetUserByEmail(c.Request.Context(), body.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := result.CheckPassword(body.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token, err := h.jwtService.GenerateToken(result.ID, result.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MapTokenToDto(token))
}

var _ authHttp.AuthHandler = &authHandler{}
