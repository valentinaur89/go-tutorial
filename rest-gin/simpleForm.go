package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var dashboard = make(map[int]Bulletin, 0)

//Using gin PostForm
/*
* How execute this program ?
* $ go run model.go simpleBind.go
* How post data ?
* curl --location --request POST 'http://localhost:8090/bulletin/ads' \
--header 'Content-Type: application/json' \
--data-raw '{"title": "amex tech",
"company":"Apple",
"product":"Pencil"}'
*
*How to get data?
* curl --location --request GET 'http://localhost:8090/bulletin/ads'
*/
func main() {
	route := gin.Default()
	route.GET("/bulletin", func(c *gin.Context) {
		fmt.Println(dashboard)
		c.JSON(http.StatusOK, dashboard)
	})
	route.POST("/bulletin", func(c *gin.Context) {
		title := c.PostForm("title")
		details := c.PostForm("details")
		author := c.PostForm("author")
		length := len(dashboard)

		dashboard[length+1] = Bulletin{
			Title:           title,
			Details:         details,
			Author:          author,
			CreateTimestamp: time.Now(),
		}
		fmt.Println(dashboard)
		c.JSON(http.StatusOK, gin.H{"title": title,
			"details": details, "author": author,
		})
	})
	route.Run(":8089")
}
