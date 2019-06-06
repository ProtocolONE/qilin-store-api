package api

import (
	jwtverifier "github.com/ProtocolONE/authone-jwt-verifier-golang"
	"github.com/ProtocolONE/qilin-store-api/pkg/api/dto"
	"github.com/ProtocolONE/qilin-store-api/pkg/api/mapper"
	"github.com/ProtocolONE/qilin-store-api/pkg/common"
	"github.com/ProtocolONE/qilin-store-api/pkg/interfaces"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

const errorAuthFailed = "Unable to authenticate user"
const userField = "user"

type AccountRouter struct {
	accountService interfaces.AccountService
}

func InitAccountRouter(group *echo.Group, accountService interfaces.AccountService) (*AccountRouter, error) {
	router := AccountRouter{accountService}

	g := group.Group("/accounts")
	g.Use(checkAuth())
	g.POST("/login", router.authorize)
	g.POST("/register", router.register)

	return &router, nil
}

func (router *AccountRouter) register(ctx echo.Context) error {
	user := ctx.Get(userField).(*jwtverifier.UserInfo)

	data := dto.RegisterAccountDTO{}
	if err := ctx.Bind(&data); err != nil {
		return common.NewServiceError(http.StatusBadRequest, errors.Wrap(err, "Binding to dto"))
	}

	if err := ctx.Validate(&data); err != nil {
		return common.NewServiceError(http.StatusUnprocessableEntity, errors.Wrap(err, "Validation failed"))
	}

	account, err := router.accountService.Register(user.UserID, data)
	if err != nil {
		return err
	}

	sess, err := session.Get("session", ctx)
	if err != nil {
		zap.L().Error("Can't get session", zap.Error(err))
	} else {
		sess.Values["email"] = account.Personal.Email
		sess.Values["user_id"] = account.ID.Hex()
		sess.Values["nickname"] = account.Account.Nickname
		err = sess.Save(ctx.Request(), ctx.Response())
		if err != nil {
			zap.L().Error("Can't save session", zap.Error(err))
		}
	}

	return ctx.JSON(http.StatusCreated, mapper.UserFromModel(account))
}

func (router *AccountRouter) authorize(ctx echo.Context) error {
	user := ctx.Get(userField).(*jwtverifier.UserInfo)

	data := dto.AuthorizeAccountDTO{}
	if err := ctx.Bind(&data); err != nil {
		return common.NewServiceError(http.StatusBadRequest, errors.Wrap(err, "Binding to dto"))
	}

	if err := ctx.Validate(&data); err != nil {
		return common.NewServiceError(http.StatusUnprocessableEntity, errors.Wrap(err, "Validation failed"))
	}

	account, err := router.accountService.Authorize(user.UserID, data)
	if err != nil {
		return err
	}

	sess, err := session.Get("session", ctx)
	if err != nil {
		zap.L().Error("Can't get session", zap.Error(err))
	} else{
		sess.Values["email"] = account.Personal.Email
		sess.Values["user_id"] = account.ID.Hex()
		sess.Values["nickname"] = account.Account.Nickname
		err = sess.Save(ctx.Request(), ctx.Response())
		if err != nil {
			zap.L().Error("Can't save session", zap.Error(err))
		}
	}

	return ctx.JSON(http.StatusOK, mapper.UserFromModel(account))
}

func checkAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			user := c.Get(userField)

			if user == nil {
				return &echo.HTTPError{
					Code:    http.StatusUnauthorized,
					Message: errorAuthFailed,
				}
			}

			return next(c)
		}
	}
}