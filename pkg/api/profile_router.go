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
	mfaService     interfaces.MfaService
}

func InitProfileRouter(group *echo.Group, service interfaces.ProfileService, mfa interfaces.MfaService) (*ProfileRouter, error) {
	router := ProfileRouter{profileService: service, mfaService: mfa}

	g := group.Group("/profiles")
	g.Use(checkAuth())

	g.GET("/me", router.getAccountInfo)
	g.PUT("/me", router.updateAccountInfo)
	g.DELETE("/me/mfa/:providerId", router.removeMfa)
	g.POST("/me/mfa/", router.addMfa)

	return &router, nil
}

func (router *ProfileRouter) addMfa(ctx echo.Context) error {
	user := ctx.Get(userField).(*jwtverifier.UserInfo)

	providerId := ctx.Param("providerId")

	if len(providerId) == 0 {
		return common.NewServiceError(http.StatusBadRequest, "Provider Id is empty")
	}

	err := router.mfaService.Add(user.UserID, providerId)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (router *ProfileRouter) removeMfa(ctx echo.Context) error {
	user := ctx.Get(userField).(*jwtverifier.UserInfo)

	providerId := ctx.Param("providerId")

	if len(providerId) == 0 {
		return common.NewServiceError(http.StatusBadRequest, "Provider Id is empty")
	}

	err := router.mfaService.Remove(user.UserID, providerId)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (router *ProfileRouter) updateAccountInfo(ctx echo.Context) error {
	user := ctx.Get(userField).(*jwtverifier.UserInfo)

	data := dto.UpdateUserDTO{}
	if err := ctx.Bind(&data); err != nil {
		return common.NewServiceError(http.StatusBadRequest, errors.Wrap(err, "Binding to dto"))
	}

	if err := ctx.Validate(&data); err != nil {
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

	data := mapper.UserFromModel(account)

	providers, err := router.mfaService.List(user.UserID)
	if err != nil {
		return err
	}

	if providers != nil {
		data.Security = &dto.UserSecurityDTO{
			MFA: mapProviders(providers),
		}
	}

	return ctx.JSON(http.StatusOK, data)
}

func mapProviders(providers []dto.MfaProviderDTO) []dto.UserMultiFactorDTO {
	var result []dto.UserMultiFactorDTO
	for _, provider := range providers {
		result = append(result, dto.UserMultiFactorDTO{
			ProviderId:   provider.ID,
			ProviderName: provider.Name,
		})
	}
	return result
}
