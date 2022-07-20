package main

import (
	"net/http"
	"time"

	// "fmt"
	"github.com/gin-gonic/gin"
)

func kst_time() string {
	utc := time.Now().UTC()
	utc = utc.In(time.FixedZone("KST", (9 * 60 * 60))) //KST = UTC + 9 hours
	kst := utc.Format("2006-01-02 15:04")              //time format
	return kst
} // Time Func

func mainpage(ctx *gin.Context) {
	kst := kst_time()
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"message": kst,
	})
}

func main() {
	router := gin.Default()
	// router.Use(mainpage)

	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")
	router.GET("/", mainpage)

	router.Run(":8080")
}
