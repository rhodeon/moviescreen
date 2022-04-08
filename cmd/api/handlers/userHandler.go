package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/domain/repository"
	"net/http"
)

type userHandler struct {
	config       common.Config
	repositories repository.Repositories
}

func NewUserHandler(config common.Config, repositories repository.Repositories) common.UserHandler {
	return &userHandler{
		config:       config,
		repositories: repositories,
	}
}

func (u userHandler) Register(ctx *gin.Context) {
	userRequest := &request.UserRequest{}
	err := parseJsonRequest(ctx, userRequest)
	if err != nil {
		return
	}

	err = validateJsonRequest(ctx, userRequest, []string{
		request.UserFieldUsername,
		request.UserFieldEmail,
		request.UserFieldPassword,
	})
	if err != nil {
		return
	}

	ctx.JSON(
		http.StatusCreated,
		response.SuccessResponse(
			http.StatusCreated,
			*userRequest,
		),
	)
}

func (u userHandler) GetByEmail(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
