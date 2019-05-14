package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/ProtocolONE/qilin-store-api/pkg/api/mapper"
	"github.com/ProtocolONE/qilin-store-api/pkg/interfaces"
)

type GameRouter struct {
	gameService interfaces.GameService
}

func InitGameRouter(group *echo.Group, gameService interfaces.GameService) (*GameRouter, error) {
	router := GameRouter{gameService: gameService}

	g := group.Group("/games")
	g.GET("/:gameId", router.getGameById)
	g.GET("/", router.getListGames)

	return &router, nil
}

func (router *GameRouter) getListGames(ctx echo.Context) error {
	panic("not implemented yet")
}

func (router *GameRouter) getGameById(ctx echo.Context) error {
	gameId := ctx.Param("gameId")

	game, err := router.gameService.GetById(gameId)
	if err != nil {
		return err
	}

	//TODO: get user language from request context
	gameDto := mapper.GameFromModel(game, "en")

	return ctx.JSON(http.StatusOK, gameDto)
}