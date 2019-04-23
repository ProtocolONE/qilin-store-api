package mapper

import (
	"super_api/pkg/api/dto"
	"super_api/pkg/model"
	"time"
)

func TagFromModel(tag model.GameTag, lng string) dto.LinkDTO {
	return dto.LinkDTO{Title: tag.Name.GetValueOrDefault(lng), Id: tag.ID }
}

func TagsFromModel(tags []model.GameTag, lng string) []dto.LinkDTO {
	var result []dto.LinkDTO
	for _, tag := range tags {
		result = append(result, TagFromModel(tag, lng))
	}
	return result
}

func LinkFromModel(link model.Link) dto.LinkDTO {
	return dto.LinkDTO{Id: link.ID, Title: link.Title}
}

func GameFromModel(game *model.Game, lng string) *dto.GameDTO {
	tags := TagsFromModel(game.Tags, lng)
	return &dto.GameDTO{
		Title:       game.Title,
		Name:        "", //TODO: узнать что ожидает фронт здесь
		Developer:   LinkFromModel(game.Developers),
		Publisher:   LinkFromModel(game.Publishers),
		Preview:     game.Previews.GetValueOrDefault(lng),
		Tags:        tags,
		ReleaseDate: game.ReleaseDate.Format(time.RFC3339),
	}
}
