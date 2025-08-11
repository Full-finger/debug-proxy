package main

import (
	"io/fs"
	"log"
	"net/http"

	"debug-proxy/internal/handler"
	"debug-proxy/internal/web"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Any("/api/request", handler.ProxyRequest)

	// 把 dist 子目录剥出来，挂载到 /static
	distRoot, _ := fs.Sub(web.DistFS, "dist")
	r.StaticFS("/static", http.FS(distRoot))

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/index.html")
	})

	log.Println("listen :8080")
	log.Fatal(r.Run(":8080"))
}
