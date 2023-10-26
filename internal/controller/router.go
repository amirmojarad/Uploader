package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router gin.IRoutes, authCtrl *Auth) {
	router.POST("/auth/signup", authCtrl.Signup)
	router.POST("/auth/login", authCtrl.Login)
}

func SetupUploadRoutes(router gin.IRoutes, middleware *Middleware, uploaderCtrl *Uploader) {
	router.Use(middleware.jwtMiddleware)

	router.POST("/file/upload", uploaderCtrl.UploadFile)
	router.GET("/file/download/:id")
	router.GET("/file/download/list", uploaderCtrl.ListFiles)
	router.GET("/file/download/all")
	router.DELETE("/file/download/:id")
}
