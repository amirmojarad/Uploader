package controller

import "github.com/gin-gonic/gin"

func setupAuthRoutes(router gin.IRoutes) {
	router.POST("/auth/signup")
	router.POST("/auth/login")
	// TODO add email verification
}

func setupUploadRoutes(router gin.IRoutes) {
	router.POST("/file/upload")
	router.GET("/file/download/:id")
	router.GET("/file/download/list")
	router.GET("/file/download/all")
	router.DELETE("/file/download/:id")
}
