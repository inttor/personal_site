package handler

import (
	"PersonalSite/backend/config"
	"PersonalSite/backend/internal/repository"
	jwtpkg "PersonalSite/backend/pkg/jwt"
	"PersonalSite/backend/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo *repository.UserRepository
	cfg      *config.Config
}

func NewAuthHandler(userRepo *repository.UserRepository, cfg *config.Config) *AuthHandler {
	return &AuthHandler{userRepo: userRepo, cfg: cfg}
}

type loginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request")
		return
	}

	user, err := h.userRepo.FindByUsername(req.Username)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	expireHours, _ := strconv.Atoi(h.cfg.JWTExpireHours)
	if expireHours <= 0 {
		expireHours = 72
	}

	token, err := jwtpkg.GenerateToken(user.ID, user.Username, h.cfg.JWTSecret, expireHours)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response.Success(c, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}
