package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"

)

var APP_MODE string
var API_PATH_BASE string
var APP_PATH string
var APP_COLOR string
var APP_IMG string

const (
		COLOR_STREAM = "#2980b9"
		COLOR_PAY = "#30336b"
		COLOR_FOOD = "#3e169d"
		COLOR_BIKE = "#020079"
		COLOR_404 = "#faa416"
		IMG_STREAM = "https://t.ctcdn.com.br/IKXEZxKaBFdLwFwFpaQpOJ-pax4=/1400x788/smart/i570349.jpeg"
		IMG_PAY = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRvr6hTX1iejidc8GUApc4gs6cZklt3HwWXrg&usqp=CAU"
		IMG_FOOD = "https://i.pinimg.com/originals/22/24/a2/2224a24c41b750dc4736648bcec5ecdc.jpg"
		IMG_BIKE = "https://i.pinimg.com/originals/b2/80/d2/b280d25cb3f3b1aa483a6b4145a220d5.jpg"
		IMG_404 = "https://i.pinimg.com/originals/c1/bb/6f/c1bb6f585059f335f1769abc1d3032b5.jpg"
)

func init(){
	
	if os.Getenv("API_PATH_BASE") !=  "" {
		API_PATH_BASE = os.Getenv("API_PATH_BASE")
	}

	if os.Getenv("APP_MODE") !=  "" {
		APP_MODE = os.Getenv("APP_MODE")
	}

	if (APP_MODE == "bike") {
		APP_COLOR = COLOR_BIKE
		APP_IMG = IMG_BIKE
	} else if (APP_MODE == "food") {
		APP_COLOR = COLOR_FOOD
		APP_IMG = IMG_FOOD
	}else if (APP_MODE == "stream") {
		APP_COLOR = COLOR_STREAM
		APP_IMG = IMG_STREAM
	}else if (APP_MODE == "pay") {
		APP_COLOR = COLOR_PAY
		APP_IMG = IMG_PAY
	} else {
		APP_COLOR = COLOR_404
		APP_IMG = IMG_404
	}

}

func main() {
	fmt.Println("simple-site")
	fmt.Printf("API_PATH_BASE %s| APP_MODE %s | APP_COLOR %s | APP_IMG %s ",API_PATH_BASE ,APP_MODE, APP_COLOR, APP_IMG)

	r := gin.Default()

	api := r.Group("/api")
	api.GET("/health", func(c *gin.Context) {
		fmt.Printf("%s api/health \n", APP_MODE)
		c.JSON(200, gin.H{
		  "health": true,
		})
	})

	r.LoadHTMLGlob("html/*.html")
	r.GET("/", func(c *gin.Context) {
		fmt.Printf("%s /{{SLASH}} \n", APP_MODE)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"color": APP_COLOR,
			"img": APP_IMG,
			"title": APP_MODE,
		})
	})

	r.GET("/pay", func(c *gin.Context) {
		fmt.Printf("%s /pay \n", APP_MODE)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"color": COLOR_PAY,
			"img": IMG_PAY,
			"title": "Pay",
		})
	})

	r.GET("/food", func(c *gin.Context) {
		fmt.Printf("%s /food \n", APP_MODE)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"color": COLOR_FOOD,
			"img": IMG_FOOD,
			"title": "Food",
		})
	})

	r.GET("/bike", func(c *gin.Context) {
		fmt.Printf("%s /bike \n", APP_MODE)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"color": COLOR_BIKE,
			"img": IMG_BIKE,
			"title": "Bike",
		})
	})

	r.GET("/stream", func(c *gin.Context) {
		fmt.Printf("%s /stream \n", APP_MODE)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"color": COLOR_STREAM,
			"img": IMG_STREAM,
			"title": "Stream",
		})
	})

	r.GET("/ops", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"color": COLOR_404,
			"img": IMG_404,
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"color": COLOR_404,
			"img": IMG_404,
		})
	})
	r.Use(static.Serve("/img", static.LocalFile("./img", true)))
	r.Run(":3000")
}