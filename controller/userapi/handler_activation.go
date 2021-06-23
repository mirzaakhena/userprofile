package userapi

import (
  "net/http"
  "userprofile/infrastructure/log"
  "userprofile/infrastructure/util"
  "userprofile/usecase/activation"

  "github.com/gin-gonic/gin"
)

// activationHandler ...
func (r *Controller) activationHandler(inputPort activation.Inport) gin.HandlerFunc {

  return func(c *gin.Context) {

    ctx := log.Context(c.Request.Context())

    req := activation.InportRequest{
      Email:           c.Param("email"),
      ActivationToken: c.Param("token"),
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
