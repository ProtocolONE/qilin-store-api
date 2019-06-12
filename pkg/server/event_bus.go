package server

import (
	"fmt"
	"github.com/ProtocolONE/qilin-common/pkg/proto"
	"github.com/ProtocolONE/qilin-store-api/pkg/interfaces"
	"github.com/ProtocolONE/qilin-store-api/pkg/model"
	"github.com/ProtocolONE/rabbitmq/pkg"
	"github.com/globalsign/mgo/bson"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"time"
)

type eventBus struct {
	provider interfaces.DatabaseProvider
	broker   *rabbitmq.Broker
	exit     chan bool
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

	w.exit = make(chan bool)
	err = w.broker.Subscribe(w.exit)
	if err != nil {
		return err
	}

	return nil
}

func (w *eventBus) Shutdown() {
	zap.L().Info("Shutdown called")
	w.exit <- true
}

func (w *eventBus) gameChanged(msg *proto.GameObject, d amqp.Delivery) (err error) {
	zap.L().Info(fmt.Sprintf("new gameChanged message `%v`", msg))

	db, err := w.provider.GetDatabase()

	if err != nil {
		zap.L().Error("Can't get db", zap.Error(err))
		return err
	}

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
		QilinID:              game.ID,
		Title:                game.Title,
		ReleaseDate:          releaseDate,
		Tags:                 mapTags(game.Tags),
		GenreMain:            mapGenre(game.GenreMain),
		GenreAddition:        mapListGenre(game.Genres),
		Requirements:         mapRequirements(game.Requirements),
		Platforms:            mapPlatforms(game.Platforms),
		Languages:            mapLanguages(game.Languages),
		Developers:           mapLink(game.Developer),
		DisplayRemainingTime: game.DisplayRemainingTime,
		FeaturesCommon:       game.Features,
		FeaturesCtrl:         game.FeaturesControl,
		Media:                mapMedia(game.Media),
		Ratings:              mapRatings(game.Ratings),
		Description:          mapLocalizedString(game.Description),
		Tagline:              mapLocalizedString(game.Tagline),
		Reviews:              mapReviews(game.Reviews),
		Publishers:           mapLink(game.Publisher),
	}
}

func mapReviews(reviews []*proto.Review) []model.GameReview {
	if reviews == nil {
		return nil
	}

	var result []model.GameReview
	for _, review := range reviews {
		result = append(result, model.GameReview{
			Link:      review.Link,
			Score:     review.Score,
			Quote:     review.Quote,
			PressName: review.PressName,
		})
	}

	return result
}

func mapRatings(ratings *proto.Ratings) *model.Ratings {
	if ratings == nil {
		return nil
	}

	return &model.Ratings{
		BBFC: mapCommonRating(ratings.BBFC),
		CERO: mapCommonRating(ratings.CERO),
		ESRB: mapCommonRating(ratings.ESRB),
		PEGI: mapCommonRating(ratings.PEGI),
		USK:  mapCommonRating(ratings.USK),
	}
}

func mapCommonRating(info *proto.RatingInfo) model.GameRating {
	if info == nil {
		return model.GameRating{}
	}

	return model.GameRating{
		AgeRestrict:         info.AgeRestrict,
		DisplayOnlineNotice: info.DisplayOnlineNotice,
		Rating:              info.Rating,
		ShowAgeRestrict:     info.ShowAgeRestrict,
	}
}

func mapMedia(media *proto.Media) *model.Media {
	if media == nil {
		return nil
	}

	return &model.Media{
		CapsuleGeneric: mapLocalizedString(media.CapsuleGeneric),
		CapsuleSmall:   mapLocalizedString(media.CapsuleSmall),
		CoverImage:     mapLocalizedString(media.CoverImage),
		CoverVideo:     mapLocalizedString(media.CapsuleGeneric),
		Friends:        mapLocalizedString(media.Friends),
		Screenshots:    mapLocalizedStringArray(media.Screenshots),
		Special:        mapLocalizedString(media.Special),
		Trailers:       mapLocalizedStringArray(media.Trailers),
	}
}

func mapLocalizedStringArray(array *proto.LocalizedStringArray) model.LocalizedStringArray {
	if array == nil {
		return model.LocalizedStringArray{}
	}

	return model.LocalizedStringArray{
		EN: array.EN,
		ES: array.ES,
		DE: array.DE,
		FR: array.FR,
		RU: array.RU,
		IT: array.IT,
		PT: array.PT,
	}
}

func mapLink(object *proto.LinkObject) model.Link {
	if object == nil {
		return model.Link{}
	}

	return model.Link{
		ID:    object.ID,
		Title: object.Title,
	}
}

func mapLanguages(languages *proto.Languages) model.GameLangs {
	if languages == nil {
		zap.L().Error("Languages is empty")
		return model.GameLangs{}
	}

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
		Voice:     language.Voice,
		Subtitles: language.Subtitles,
		Interface: language.Interface,
	}
}

func mapPlatforms(platforms *proto.Platforms) model.Platforms {
	if platforms == nil {
		zap.L().Error("Platforms is empty")
		return model.Platforms{}
	}

	return model.Platforms{
		Windows: platforms.Windows,
		Linux:   platforms.Linux,
		MacOs:   platforms.MacOs,
	}
}

func mapRequirements(requirements *proto.Requirements) model.GameRequirements {
	if requirements == nil {
		zap.L().Error("GameRequirements is empty")
		return model.GameRequirements{}
	}

	return model.GameRequirements{
		MacOs:   mapPlatformReq(requirements.MacOs),
		Linux:   mapPlatformReq(requirements.Linux),
		Windows: mapPlatformReq(requirements.Windows),
	}
}

func mapPlatformReq(requirements *proto.PlatformRequirements) *model.PlatformRequirements {
	if requirements == nil {
		return nil
	}

	return &model.PlatformRequirements{
		Recommended: mapMachineReq(requirements.Recommended),
		Minimal:     mapMachineReq(requirements.Minimal),
	}
}

func mapMachineReq(requirements *proto.MachineRequirements) model.MachineRequirements {
	if requirements == nil {
		zap.L().Error("MachineRequirements is empty")
		return model.MachineRequirements{}
	}

	return model.MachineRequirements{
		StorageDimension: requirements.StorageDimension,
		Storage:          requirements.Storage,
		Sound:            requirements.Sound,
		Ram:              requirements.Ram,
		RamDimension:     requirements.RamDimension,
		Processor:        requirements.Processor,
		Other:            requirements.Other,
		Graphics:         requirements.Graphics,
		System:           requirements.System,
	}
}

func mapListGenre(objects []*proto.TagObject) []model.GameGenre {
	if objects == nil {
		return nil
	}

	var result []model.GameGenre
	for _, object := range objects {
		result = append(result, mapGenre(object))
	}
	return result
}

func mapGenre(genre *proto.TagObject) model.GameGenre {
	if genre == nil {
		return model.GameGenre{}
	}

	return model.GameGenre{
		Name: mapLocalizedString(genre.Name),
		ID:   genre.ID,
	}
}

func mapTags(tags []*proto.TagObject) []model.GameTag {
	if tags == nil {
		return nil
	}

	var res []model.GameTag
	for _, tag := range tags {
		res = append(res, mapTag(tag))
	}
	return res
}

func mapTag(tag *proto.TagObject) model.GameTag {
	if tag == nil {
		return model.GameTag{}
	}
	return model.GameTag{ID: tag.ID, Name: mapLocalizedString(tag.Name)}
}

func mapLocalizedString(s *proto.LocalizedString) model.LocalizedString {
	if s == nil {
		return model.LocalizedString{}
	}

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
