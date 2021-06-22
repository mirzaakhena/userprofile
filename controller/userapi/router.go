package userapi

import (
	"userprofile/infrastructure/token"
	"userprofile/usecase/activation"
	"userprofile/usecase/loginuser"
	"userprofile/usecase/registeruser"
	"userprofile/usecase/showalluser"
	"userprofile/usecase/showuser"
	"userprofile/usecase/updateuser"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserToken          *token.JWTToken
	Router             gin.IRouter
	ActivationInport   activation.Inport
	LoginUserInport    loginuser.Inport
	RegisterUserInport registeruser.Inport
	ShowAllUserInport  showalluser.Inport
	ShowUserInport     showuser.Inport
	UpdateUserInport   updateuser.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/activation", r.authorized(), r.activationHandler(r.ActivationInport))
	r.Router.POST("/loginuser", r.authorized(), r.loginUserHandler(r.LoginUserInport))
	r.Router.POST("/registeruser", r.authorized(), r.registerUserHandler(r.RegisterUserInport))
	r.Router.GET("/showalluser", r.authorized(), r.showAllUserHandler(r.ShowAllUserInport))
	r.Router.GET("/showuser", r.authorized(), r.showUserHandler(r.ShowUserInport))
	r.Router.PUT("/updateuser", r.authorized(), r.updateUserHandler(r.UpdateUserInport))
}
