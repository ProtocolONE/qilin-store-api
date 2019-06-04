package api

import (
	jwtverifier "github.com/ProtocolONE/authone-jwt-verifier-golang"
	"github.com/ProtocolONE/qilin-store-api/pkg/api/dto"
	"github.com/ProtocolONE/qilin-store-api/pkg/api/mapper"
	"github.com/ProtocolONE/qilin-store-api/pkg/common"
	"github.com/ProtocolONE/qilin-store-api/pkg/interfaces"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

type ProfileRouter struct {
	profileService interfaces.ProfileService
}

func InitProfileRouter(group *echo.Group, service interfaces.ProfileService) (*ProfileRouter, error) {
	router := ProfileRouter{profileService: service}

	g := group.Group("/profiles")
	g.Use(checkAuth())

	g.GET("/me", router.getAccountInfo)
	g.PUT("/me", router.updateAccountInfo)

	return &router, nil
}

func (router *ProfileRouter) updateAccountInfo(ctx echo.Context) error {
	user := ctx.Get(userField).(*jwtverifier.UserInfo)

	data := dto.UpdateUserDTO{}
	if err := ctx.Bind(data); err != nil {
		return common.NewServiceError(http.StatusBadRequest, errors.Wrap(err, "Binding to dto"))
	}

	if err := ctx.Validate(data); err != nil {
		return common.NewServiceError(http.StatusUnprocessableEntity, errors.Wrap(err, "Validation failed"))
	}

	account, err := router.profileService.UpdateAccount(user.UserID, data)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, mapper.UserFromModel(account))
}

func (router *ProfileRouter) getAccountInfo(ctx echo.Context) error {
	user := ctx.Get(userField).(*jwtverifier.UserInfo)

	account, err := router.profileService.GetAccount(user.UserID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, mapper.UserFromModel(account))
}