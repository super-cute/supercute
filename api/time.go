package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) Time(c *gin.Context) {
	execute := func() error {
		return nil
	}

	err := execute()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
	})
}
