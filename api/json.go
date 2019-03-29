package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (api *Api) UnMarshal(c *gin.Context) {
	execute := func() (string, error) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return "", err
		}

		tmp := make(map[string]interface{})

		if err := json.Unmarshal(body, &tmp); err != nil {
			return "", err
		}

		nb, err := json.Marshal(tmp)
		if err != nil {
			return "", err
		}
		res := string(nb)
		res = strings.Replace(res, "\n", "", -1)
		res = strings.Replace(res, "\"", "\\\"", -1)

		return res, nil
	}

	res, err := execute()
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
		"data":    res,
	})
}

func (api *Api) JsonUnescape(c *gin.Context) {
	execute := func() (interface{}, error) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return nil, err
		}

		res := string(body)
		res = strings.Replace(res, "\\\"", "\"", -1)

		tmp := make(map[string]interface{})

		if err := json.Unmarshal([]byte(res), &tmp); err != nil {
			return nil, err
		}

		return &tmp, nil
	}

	res, err := execute()
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
		"data":    res,
	})
}
