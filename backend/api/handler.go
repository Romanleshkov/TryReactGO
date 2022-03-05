package api

import (
	"backend/ansibleVault"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/encrypt", h.encrypt)
		api.POST("/decrypt", h.decrypt)
	}

	return router
}

func (h *Handler) encrypt(c *gin.Context) {
	var input Request
	status := "OK"

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	content, err := ansibleVault.Encrypt(input.Content, input.Password)
	if err != nil {
		status = "Error of encryption"
	}
	c.JSONP(http.StatusOK, map[string]interface{}{
		"content": content,
		"status":  status,
	})
}

func (h *Handler) decrypt(c *gin.Context) {
	var input Request

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	status := "OK"

	content, err := ansibleVault.Decrypt(input.Content, input.Password)
	if err != nil {
		status = "Error of decryption"
	}
	c.JSONP(http.StatusOK, map[string]interface{}{
		"content": content,
		"status":  status,
	})
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, struct {
		Message string `json:"message"`
	}{message})
}
