package server

func Start() {
	router := getRouter()
	router.Run(":8080")
}
