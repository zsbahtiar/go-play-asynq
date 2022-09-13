package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zsbahtiar/go-play-asynq/client/module/user"
	"net/http"
)

type userHandler struct {
	userSvc user.Service
}

type UserHandler interface {
	CreateUsersCsv(ctx *gin.Context)
}

func NewUserHandler(userSvc user.Service) UserHandler {
	return &userHandler{userSvc}
}

func (u *userHandler) CreateUsersCsv(ctx *gin.Context) {
	var req user.CreateUsersCsvRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"isSuccess": false,
			"message":   err.Error(),
		})
		return
	}
	if req.FileURL == "" {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"isSuccess": false,
			"message":   "fileURL is required",
		})
		return
	}

	if err = u.userSvc.CreateUsersCsv(ctx, req.FileURL); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"isSuccess": false,
			"message":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"isSuccess": true,
		"message":   "success create users",
	})
	return
}
