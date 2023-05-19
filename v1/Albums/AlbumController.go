package albums

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	service AlbumService
}

func NewAlbumController(service AlbumService) *AlbumController {
	return &AlbumController{
		service: service,
	}
}

func (ac *AlbumController) GetAlbums(c *gin.Context) {
	albums, err := ac.service.GetAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}
