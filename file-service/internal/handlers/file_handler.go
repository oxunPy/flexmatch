package handlers

import (
	"file-service/internal/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type FileUploadResponse struct {
	Success bool             `json:"success"`
	File    *models.Datafile `json:"file"`
}

func (rest *RestController) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file not found"})
		return
	}
	defer file.Close()

	contentType := header.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	content := make([]byte, header.Size)
	if _, err := file.Read(content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fayl o'qib bo'lmadi"})
		return
	}

	datafile, err := rest.FileService.Upload(c.Request.Context(), header.Filename, contentType, 1, content)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, FileUploadResponse{
		Success: true,
		File:    datafile,
	})
}

func (rest *RestController) GetInfo(c *gin.Context) {
	fileID := c.Param("id")

	f, err := rest.FileService.GetFile(c.Request.Context(), fileID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"file_id":      f.ID,
		"name":         f.Name,
		"content_type": f.ContentType,
		"size":         f.Size,
		"url":          f.URL,
		"uploaded_by":  f.UploadedBy,
		"created_at":   f.CreatedAt,
	})
}

func (rest *RestController) Download(c *gin.Context) {
	fileID := c.Param("id")

	f, content, err := rest.FileService.GetContent(c.Request.Context(), fileID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filepath.Base(f.Name))
	c.Data(http.StatusOK, f.ContentType, content)
}

func (rest *RestController) Delete(c *gin.Context) {
	fileID := c.Param("id")

	if err := rest.FileService.Delete(c.Request.Context(), fileID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
