package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func Error(c *gin.Context, message string, code int) {
	c.Header("HX-Retarget", "#errors")
	c.HTML(code, "error.html", message)
}
