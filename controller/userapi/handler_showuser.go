package userapi

import (
	"net/http"
	"userprofile/infrastructure/log"
	"userprofile/infrastructure/util"
	"userprofile/usecase/showuser"

	"github.com/gin-gonic/gin"
)

// showUserHandler ...
func (r *Controller) showUserHandler(inputPort showuser.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.Context(c.Request.Context())

		var req showuser.InportRequest
		req.UserID = c.Param("userID")

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
