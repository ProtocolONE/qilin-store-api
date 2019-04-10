package mapper

import (
	"super_api/pkg/api/dto"
	"super_api/pkg/model"
	"time"
)

func TagFromModel(tag model.GameTag, lng string) string {
	return tag.Name.GetValueOrDefault(lng)
}

func TagsFromModel(tags []model.GameTag, lng string) []string {
	var result []string
	for _, tag := range tags {
		result = append(result, TagFromModel(tag, lng))
	}
	return result
}

func GameFromModel(game *model.Game, lng string) *dto.GameDTO {
	tags := TagsFromModel(game.Tags, lng)
	return &dto.GameDTO{
		Title:       game.Title,
		Name:        "", //TODO: узнать что ожидает фронт здесь
		Developer:   game.Developers,
		Publisher:   game.Publishers,
		Preview:     game.Previews.GetValueOrDefault(lng),
		Tags:        tags,
		ReleaseDate: game.ReleaseDate.Format(time.RFC3339),
		//TODO
	}
}
