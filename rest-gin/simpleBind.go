package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//Using gin Json Request
/*
* How to execute this program ?
* $ go run model.go simpleBind.go
* How post data ?
* curl --location --request POST 'http://localhost:8090/bulletin' \
--header 'Content-Type: application/json' \
--form 'title=" amex tech"' \
--form 'company="Apple"' \
--form 'product="Pencil"'
*
*How to get data?
* curl --location --request GET 'http://localhost:8090/bulletin'
*/
func main() {

	var adsMap = make(map[int]Ads, 0)
	route := gin.Default()
	route.GET("/bulletin/ads", func(c *gin.Context) {
		c.JSON(http.StatusOK, adsMap)
	})
	route.POST("/bulletin/ads", func(c *gin.Context) {
		var ad Ads
		c.BindJSON(&ad)
		l := len(adsMap)
		adsMap[l+1] = Ads{
			Title:           ad.Title,
			Company:         ad.Company,
			Product:         ad.Product,
			CreateTimestamp: time.Now(),
		}
		c.JSON(http.StatusOK, gin.H{"title": ad.Title, "company:": ad.Company, "product": ad.Product})
	})
	route.Run(":8090")
}
