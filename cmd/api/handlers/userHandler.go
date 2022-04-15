package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/domain/repository"
	"github.com/rhodeon/moviescreen/internal/mailer"
	"github.com/rhodeon/moviescreen/internal/validator"
	"github.com/rhodeon/prettylog"
	"net/http"
	"sync"
	"time"
)

type userHandler struct {
	config       common.Config
	repositories repository.Repositories
	backgroundWg *sync.WaitGroup
}

func NewUserHandler(config common.Config, repositories repository.Repositories, waitGroup *sync.WaitGroup) common.UserHandler {
	return &userHandler{
		config:       config,
		repositories: repositories,
		backgroundWg: waitGroup,
	}
}

func (u userHandler) Register(ctx *gin.Context) {
	// parse request body
	userRequest := &request.UserRequest{}
	err := parseJsonRequest(ctx, userRequest)
	if err != nil {
		return
	}

	// validate request body
	err = validateJsonRequest(ctx, userRequest, []string{
		request.UserFieldUsername,
		request.UserFieldEmail,
		request.UserFieldPassword,
	})
	if err != nil {
		return
	}

	// map request body to user body for further operations
	user, err := userRequest.ToModel()
	if err != nil {
		HandleInternalServerError(ctx, err)
		return
	}

	// attempt to register user
	err = u.repositories.Users.Register(&user)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrDuplicateUsername):
			v := validator.New("user")
			v.AddError(request.UserFieldUsername, "this username is already taken")
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				response.UnprocessableEntityError(v),
			)

		case errors.Is(err, repository.ErrDuplicateEmail):
			v := validator.New("user")
			v.AddError(request.UserFieldEmail, "a user with this email address already exists")
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				response.UnprocessableEntityError(v),
			)

		default:
			HandleInternalServerError(ctx, err)
		}

		return
	}

	// generate activation token with a lifetime of 2 days
	token, err := u.repositories.Tokens.New(user.Id, models.ScopeActivation, 2*24*time.Hour)
	if err != nil {
		HandleInternalServerError(ctx, err)
	}

	// send welcome email to user in the background
	common.Background(u.backgroundWg, func() {
		smtp := u.config.Smtp
		mail := mailer.New(smtp.Host, smtp.Port, smtp.User, smtp.Password, smtp.Sender)

		err = mail.Send(user.Email, "user_welcome.gotmpl", struct {
			Username        string
			ActivationToken string
		}{
			Username:        user.Username,
			ActivationToken: token.PlainText,
		})

		if err != nil {
			prettylog.ErrorF("Welcome mail: %v", err)
			return
		}
	})

	// return new user details
	ctx.JSON(
		http.StatusCreated,
		response.SuccessResponse(
			http.StatusCreated,
			user.ToResponse(),
		),
	)
}

func (u userHandler) Activate(ctx *gin.Context) {
	// parse request body
	req := &request.UserActivationRequest{}
	err := parseJsonRequest(ctx, req)
	if err != nil {
		return
	}

	// validate request token
	err = validateJsonRequest(ctx, req, []string{request.UserActivationFieldToken})
	if err != nil {
		return
	}

	// get user associated with token
	user, err := u.repositories.Users.GetByToken(*req.Token, models.ScopeActivation)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			v := validator.New("activate")
			v.AddError(request.UserActivationFieldToken, "invalid or expired activation token")
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				response.UnprocessableEntityError(v),
			)

		default:
			HandleInternalServerError(ctx, err)
		}

		return
	}

	// activate user
	user.Activated = true

	// save user activated status
	err = u.repositories.Users.Update(&user)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrEditConflict):
			NewErrorHandler().EditConflict(ctx)

		default:
			HandleInternalServerError(ctx, err)
		}

		return
	}

	// delete used token
	err = u.repositories.Tokens.DeleteAllForUser(user.Id, models.ScopeActivation)
	if err != nil {
		HandleInternalServerError(ctx, err)
		return
	}

	// return updated user response
	ctx.JSON(
		http.StatusOK,
		response.SuccessResponse(
			http.StatusOK,
			user.ToResponse(),
		),
	)
}

func (u userHandler) Authenticate(ctx *gin.Context) {
	req := &request.UserRequest{}
	err := parseJsonRequest(ctx, req)
	if err != nil {
		return
	}

	err = validateJsonRequest(ctx, req, []string{request.UserFieldEmail, request.UserFieldPassword})
	if err != nil {
		return
	}

	// retrieve user data via email
	user, err := u.repositories.Users.GetByEmail(*req.Email)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			NewErrorHandler().InvalidCredentials(ctx)

		default:
			HandleInternalServerError(ctx, err)
		}

		return
	}

	// return unauthorized error if the user is not activated
	if !user.Activated {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.ErrorResponse(
				http.StatusUnauthorized,
				response.GenericError("user not activated"),
			),
		)
		return
	}

	// confirm password
	valid, err := user.Password.Matches(*req.Password)
	if err != nil {
		HandleInternalServerError(ctx, err)
		return
	}
	if !valid {
		NewErrorHandler().InvalidCredentials(ctx)
		return
	}

	// generate new authentication token with a lifetime of 1 day
	token, err := u.repositories.Tokens.New(user.Id, models.ScopeAuthentication, 1*24*time.Hour)
	if err != nil {
		HandleInternalServerError(ctx, err)
		return
	}

	// return token as response
	ctx.JSON(
		http.StatusCreated,
		response.SuccessResponse(
			http.StatusCreated,
			token.ToResponse(),
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
