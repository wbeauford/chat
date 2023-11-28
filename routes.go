package main

func initalizeRoutes() {
	router.GET("/", renderIndex)
	router.GET("/getRooms", getRooms)
	router.POST("/createRoom", createRoom)
	router.POST("/joinRoom", joinRoom)
	router.GET("/ws", wsHandler)}
