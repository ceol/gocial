package server

import api "github.com/ceol/gocial/internal/api/routes"

func Start() {
	router := NewRouter()

	apiGroup := router.Group("/api")
	api.AddRoutes(apiGroup)

	router.Run()
}
