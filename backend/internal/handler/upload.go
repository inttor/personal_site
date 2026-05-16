package handler

import (
	"PersonalSite/backend/config"
	"PersonalSite/backend/pkg/response"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	cfg *config.Config
}

func NewUploadHandler(cfg *config.Config) *UploadHandler {
	return &UploadHandler{cfg: cfg}
}

func (h *UploadHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "No file provided")
		return
	}

	// Validate file size (max 10MB)
	if file.Size > 10<<20 {
		response.Error(c, http.StatusBadRequest, "File too large (max 10MB)")
		return
	}

	// Validate file extension
	ext := filepath.Ext(file.Filename)
	allowed := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
		".webp": true, ".svg": true, ".bmp": true,
	}
	if !allowed[ext] {
		response.Error(c, http.StatusBadRequest, "File type not allowed")
		return
	}

	// Create upload dir if not exists
	uploadDir := h.cfg.UploadDir
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create upload directory")
		return
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dst := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to save file")
		return
	}

	response.Success(c, gin.H{
		"url": "/uploads/" + filename,
	})
}
