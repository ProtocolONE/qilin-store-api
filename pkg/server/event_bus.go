package server

import (
	"fmt"
	"github.com/ProtocolONE/rabbitmq/pkg"
	"github.com/globalsign/mgo/bson"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"super_api/pkg/interfaces"
	"super_api/pkg/model"
	"github.com/ProtocolONE/qilin-common/pkg/proto"
	"time"
)

type eventBus struct {
	provider interfaces.DatabaseProvider
	broker *rabbitmq.Broker
	exit 	chan bool
}

func NewEventBus(provider interfaces.DatabaseProvider, host string) (interfaces.EventBus, error) {
	zap.L().Info("Creating broker for AMPQ")

	broker, err := rabbitmq.NewBroker(host)
	if err != nil {
		return nil, err
	}

	return &eventBus{provider: provider, broker: broker}, nil
}

func (w *eventBus) StartListen() error {
	err := w.broker.RegisterSubscriber("game_changed", w.gameChanged)
	if err != nil {
		return err
	}

	w.exit= make(chan bool)
	err = w.broker.Subscribe(w.exit)
	if err != nil {
		return err
	}

	return nil
}

func (w *eventBus) Shutdown() {
	w.exit <- true
}

func (w *eventBus) gameChanged(msg *proto.GameObject, d amqp.Delivery) (err error) {
	zap.L().Debug(fmt.Sprintf("new gameChanged message `%v`", msg))

	db, err := w.provider.GetDatabase()

	game := mapGame(msg)
	_, err = db.C("games").Upsert(bson.M{"qilin_id": game.QilinID}, game)

	if err != nil {
		zap.L().Error("Can't save game", zap.Error(err))
	}

	return err
}

func mapGame(game *proto.GameObject) *model.Game {
	releaseDate, _ := time.Parse(time.RFC3339, game.ReleaseDate)
	return &model.Game{
		QilinID:     game.ID,
		Title:       game.Title,
		ReleaseDate: releaseDate,
		Tags:        mapTags(game.Tags),
		GenreMain:   mapGenre(game.GenreMain),
		GenreAddition: mapListGenre(game.Genres),
		Requirements: mapRequirements(game.Requirements),
		Platforms: mapPlatforms(game.Platforms),
		Languages: mapLanguages(game.Languages),
		Developers: mapLink(game.Developer),
		DisplayRemainingTime: game.DisplayRemainingTime,
		FeaturesCommon: game.Features,
		FeaturesCtrl: game.FeaturesControl,
	}
}

func mapLink(object *proto.LinkObject) model.Link {
	return model.Link{
		ID: object.ID,
		Title: object.Title,
	}
}

func mapLanguages(languages *proto.Languages) model.GameLangs {
	return model.GameLangs{
		EN: mapLang(languages.EN),
		IT: mapLang(languages.IT),
		PT: mapLang(languages.PT),
		FR: mapLang(languages.FR),
		ES: mapLang(languages.ES),
		DE: mapLang(languages.DE),
		RU: mapLang(languages.RU),
	}
}

func mapLang(language *proto.Language) model.Langs {
	if language == nil {
		return model.Langs{}
	}

	return model.Langs{
		Voice: language.Voice,
		Subtitles: language.Subtitles,
		Interface: language.Interface,
	}
}

func mapPlatforms(platforms *proto.Platforms) model.Platforms {
	return model.Platforms{
		Windows: platforms.Windows,
		Linux: platforms.Linux,
		MacOs: platforms.MacOs,
	}
}

func mapRequirements(requirements *proto.Requirements) model.GameRequirements {
	return model.GameRequirements{
		MacOs: mapPlatformReq(requirements.MacOs),
		Linux: mapPlatformReq(requirements.Linux),
		Windows: mapPlatformReq(requirements.Windows),
	}
}

func mapPlatformReq(requirements *proto.PlatformRequirements) model.PlatformRequirements {
	return model.PlatformRequirements{
		Recommended: mapMachineReq(requirements.Recommended),
		Minimal: mapMachineReq(requirements.Minimal),
	}
}

func mapMachineReq(requirements *proto.MachineRequirements) model.MachineRequirements {
	return model.MachineRequirements{
		StorageDimension: requirements.StorageDimension,
		Storage: requirements.Storage,
		Sound: requirements.Sound,
		Ram: requirements.Ram,
		RamDimension: requirements.RamDimension,
		Processor: requirements.Processor,
		Other: requirements.Other,
		Graphics: requirements.Graphics,
		System: requirements.System,
	}
}

func mapListGenre(objects []*proto.TagObject) []model.GameGenre {
	var result []model.GameGenre
	for _, object := range objects {
		result = append(result, mapGenre(object))
	}
	return result
}

func mapGenre(genre *proto.TagObject) model.GameGenre {
	return model.GameGenre{
		Name: mapLocalizedString(genre.Name),
		ID:   genre.ID,
	}
}

func mapTags(tags []*proto.TagObject) []model.GameTag {
	var res []model.GameTag
	for _, tag := range tags {
		res = append(res, mapTag(tag))
	}
	return res
}

func mapTag(tag *proto.TagObject) model.GameTag {
	return model.GameTag{ID: tag.ID, Name: mapLocalizedString(tag.Name)}
}

func mapLocalizedString(s *proto.LocalizedString) model.LocalizedString {
	return model.LocalizedString{
		EN: s.EN,
		RU: s.RU,
		PT: s.PT,
		IT: s.IT,
		FR: s.FR,
		ES: s.ES,
		DE: s.DE,
	}
}
