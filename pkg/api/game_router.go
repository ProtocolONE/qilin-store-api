package api

import (
	"github.com/ProtocolONE/qilin-store-api/pkg/api/dto"
	"github.com/ProtocolONE/qilin-store-api/pkg/api/mapper"
	"github.com/ProtocolONE/qilin-store-api/pkg/interfaces"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GameRouter struct {
	gameService interfaces.GameService
}

func InitGameRouter(group *echo.Group, gameService interfaces.GameService) (*GameRouter, error) {
	router := GameRouter{gameService: gameService}

	g := group.Group("/games")
	g.GET("/", router.getListGames)
	g.GET("/:gameId", router.getGameById)

	return &router, nil
}

func (router *GameRouter) getListGames(ctx echo.Context) error {
	games, err := router.gameService.GetListGames("", 0, 0, "")
	if err != nil {
		return err
	}

	var result []*dto.GameDTO
	for _, game := range games {
		result = append(result, mapper.GameFromModel(&game, "en"))
	}

	return ctx.JSON(http.StatusOK, games)
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
