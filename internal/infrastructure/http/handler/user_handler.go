package handler

import (
	"net/http"
	userApp "userman/internal/application/user"
	"userman/internal/infrastructure/jwt"
	userHttp "userman/internal/interface/http/user"
	"userman/internal/interface/http/user/dto"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService userApp.UserService
}

func NewUserHandler(userService userApp.UserService) userHttp.UserHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	value := c.MustGet("claim")
	claims, _ := value.(*jwt.Claims)
	if !claims.IsAdmin() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var body dto.CreateUserBodyRequest
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

	c.JSON(http.StatusCreated, dto.MapDomainToDto(result))
}

func (h *userHandler) GetUser(c *gin.Context) {
	var uri dto.GetUserQueryRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.GetUserByID(c.Request.Context(), uri.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MapDomainToDto(result))
}

func (h *userHandler) GetAllUser(c *gin.Context) {
	var query dto.GetAllUserQueryRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := query.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.GetAllUser(c.Request.Context(), &userApp.GetAllUserInput{
		NextCursor: query.NextToken,
		PrevCursor: query.PrevToken,
		Limit:      query.Limit,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MapCursorDomainToDto(result))
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var uri dto.UpdateUserUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body dto.UpdateUserBodyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.UpdateUserByID(c.Request.Context(), uri.ID, body.Name, body.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MapDomainToDto(result))
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	value := c.MustGet("claims")
	claims, _ := value.(*jwt.Claims)
	if !claims.IsAdmin() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var uri dto.DeleteUserUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.DeleteUserByID(c.Request.Context(), uri.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MapDomainToDto(result))
}

var _ userHttp.UserHandler = &userHandler{}
