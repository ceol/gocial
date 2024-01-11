package api

import "github.com/gin-gonic/gin"

func AddRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	addUserRoutes(v1)
}
