package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getuserById/:userId", FindUserByID)
	r.GET("/getuserByEmail/:userEmail", FindUserByEmail)
	r.POST("/createUser", CreateUser)
	r.PUT("/updateUser/:userId", UpdateUser)
	r.DELETE("/deleteUser/:userId", DeleteUser)
}
