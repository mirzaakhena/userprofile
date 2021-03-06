package userapi

import (
	"net/http"
	"userprofile/application/apperror"
	"userprofile/infrastructure/log"
	"userprofile/infrastructure/util"
	"userprofile/usecase/loginuser"

	"github.com/gin-gonic/gin"
)

// loginUserHandler ...
func (r *Controller) loginUserHandler(inputPort loginuser.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.Context(c.Request.Context())

		var req loginuser.InportRequest
		if err := c.BindJSON(&req); err != nil {
			newErr := apperror.FailUnmarshalResponseBodyError
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(newErr))
			return
		}

		log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(err))
			return
		}

		log.Info(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, NewSuccessResponse(res))

	}
}
