package mapper

import (
	"fmt"
	"github.com/ProtocolONE/qilin-store-api/pkg/api/dto"
	"github.com/ProtocolONE/qilin-store-api/pkg/model"
	"time"
)

func TagFromModel(tag model.GameTag, lng string) dto.LinkDTO {
	return dto.LinkDTO{Title: tag.Name.GetValueOrDefault(lng), Id: tag.ID}
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
		Title:        game.Title,
		Name:         game.Title,
		Developer:    LinkFromModel(game.Developers),
		Publisher:    LinkFromModel(game.Publishers),
		Preview:      game.Previews.GetValueOrDefault(lng),
		Tags:         tags,
		ReleaseDate:  game.ReleaseDate.Format(time.RFC3339),
		Platforms:    PlatformsFromModel(game.Platforms),
		Requirements: RequirementsFromModel(game.Requirements, game.Languages),
	}
}

func RequirementsFromModel(requirements model.GameRequirements, langs model.GameLangs) dto.GameRequirementsDTO {
	req := dto.GameRequirementsDTO{}

	req.Systems = map[string]dto.SystemsDTO{}


	if requirements.MacOs != nil {
		req.Systems["mac_os"] = dto.SystemsDTO{
			Minimal:     SystemFromModel(requirements.MacOs.Minimal),
			Recommended: SystemFromModel(requirements.MacOs.Recommended),
		}
	}

	if requirements.Windows != nil {
		req.Systems["windows"] = dto.SystemsDTO{
			Minimal:     SystemFromModel(requirements.Windows.Minimal),
			Recommended: SystemFromModel(requirements.Windows.Recommended),
		}
	}

	if requirements.Linux != nil {
		req.Systems["linux"] = dto.SystemsDTO{
			Minimal:     SystemFromModel(requirements.Linux.Minimal),
			Recommended: SystemFromModel(requirements.Linux.Recommended),
		}
	}

	req.Languages.Audio = mapAutioFromLangs(langs)
	req.Languages.Text = mapTextFromLangs(langs)

	return req
}

func mapTextFromLangs(langs model.GameLangs) []string {
	var result []string
	result = appendTextIfNeeded(result, langs.IT, "it")
	result = appendTextIfNeeded(result, langs.RU, "ru")
	result = appendTextIfNeeded(result, langs.FR, "fr")
	result = appendTextIfNeeded(result, langs.DE, "de")
	result = appendTextIfNeeded(result, langs.EN, "en")
	result = appendTextIfNeeded(result, langs.ES, "es")
	result = appendTextIfNeeded(result, langs.PT, "pt")
	return result
}

func appendTextIfNeeded(arr []string, lang model.Langs, name string) []string {
	if lang.Interface {
		arr = append(arr, name)
	}
	return arr
}

func mapAutioFromLangs(langs model.GameLangs) []string {
	var result []string
	result = appendSoundIfNeeded(result, langs.IT, "it")
	result = appendSoundIfNeeded(result, langs.RU, "ru")
	result = appendSoundIfNeeded(result, langs.FR, "fr")
	result = appendSoundIfNeeded(result, langs.DE, "de")
	result = appendSoundIfNeeded(result, langs.EN, "en")
	result = appendSoundIfNeeded(result, langs.ES, "es")
	result = appendSoundIfNeeded(result, langs.PT, "pt")
	return result
}

func appendSoundIfNeeded(arr []string, lang model.Langs, name string) []string {
	if lang.Voice {
		arr = append(arr, name)
	}
	return arr
}

func SystemFromModel(requirements model.MachineRequirements) dto.RequirementsDTO {
	return dto.RequirementsDTO{
		CPU:       requirements.Processor,
		GPU:       requirements.Graphics,
		OS:        requirements.System,
		RAM:       fmt.Sprintf("%d %s", requirements.Ram, requirements.RamDimension),
		DiskSpace: fmt.Sprintf("%d %s", requirements.Storage, requirements.StorageDimension),
	}
}

func PlatformsFromModel(platforms model.Platforms) []string {
	var result []string
	if platforms.Linux {
		result = append(result, "linux")
	}
	if platforms.Windows {
		result = append(result, "windows")
	}
	if platforms.MacOs {
		result = append(result, "mac_os")
	}
	return result
}
