package routes

import (
	controllers "github.com/didiyudha/go-mongodb-restful/controllers"
	"github.com/julienschmidt/httprouter"
)

var usrCtrl *controllers.UsersController

func init() {
	usrCtrl = controllers.NewUserController()
}

// NewRouter return a pointer router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", controllers.Index)
	router.POST("/users", usrCtrl.SaveUser)
	router.GET("/users", usrCtrl.GetUsers)
	router.GET("/users/:id", usrCtrl.FindUser)
	router.PUT("/users/:id", usrCtrl.UpdateUser)
	router.DELETE("/users/:id", usrCtrl.DeleteUser)
	return router
}
