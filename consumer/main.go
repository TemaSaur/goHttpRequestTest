package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Format struct {
	Key1   int    `json:"key1"`
	Key2   string `json:"key2"`
	Values []int  `json:"values"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		res, err := http.Get("http://localhost:8000/")
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		var data []Format
		json.Unmarshal(body, &data)

		c.JSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
