package model

import "time"

type (
	Descriptor struct {
		Title  LocalizedString
		System string
	}

	Game struct {
		QilinID              string           `json:"qilin_id"`
		Title                string           `json:"title"`
		Developers           Link             `json:"developers"`
		Publishers           Link             `json:"publishers"`
		ReleaseDate          time.Time        `json:"release_date"`
		DisplayRemainingTime bool             `json:"display_remaining_time"`
		AchievementOnProd    bool             `json:"achievement_on_prod"`
		FeaturesCommon       []string         `json:"features_common"`
		FeaturesCtrl         string           `json:"features_ctrl"`
		Platforms            Platforms        `json:"platforms"`
		Requirements         GameRequirements `json:"requirements"`
		Languages            GameLangs        `json:"languages"`
		GenreMain            GameGenre        `json:"genre_main"`
		GenreAddition        []GameGenre      `json:"genre_addition"`
		Tags                 []GameTag        `json:"tags"`
		Previews             LocalizedString  `json:"previews"`
	}

	Link struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}

	GameGenre struct {
		ID   string `json:"id"`
		Name LocalizedString
	}

	GameTag struct {
		ID   string `json:"id"`
		Name LocalizedString
	}

	GameDescription struct {
		Tagline               LocalizedString `json:"tagline"`
		Description           LocalizedString `json:"description"`
		Reviews               GameReviews     `json:"reviews"`
		AdditionalDescription string          `json:"additional_description"`
		GameSite              string          `json:"game_site"`
		Socials               Socials         `json:"socials"`
	}

	MachineRequirements struct {
		System           string `json:"system"`
		Processor        string `json:"processor"`
		Graphics         string `json:"graphics"`
		Sound            string `json:"sound"`
		Ram              int32    `json:"ram"`
		RamDimension     string `json:"ram_dimension"`
		Storage          int32    `json:"storage"`
		StorageDimension string `json:"storage_dimension"`
		Other            string `json:"other"`
	}

	PlatformRequirements struct {
		Minimal     MachineRequirements `json:"minimal"`
		Recommended MachineRequirements `json:"recommended"`
	}

	GameRequirements struct {
		Windows PlatformRequirements `json:"windows"`
		MacOs   PlatformRequirements `json:"mac_os"`
		Linux   PlatformRequirements `json:"linux"`
	}

	Platforms struct {
		Windows bool `json:"windows"`
		MacOs   bool `json:"mac_os"`
		Linux   bool `json:"linux"`
	}

	Langs struct {
		Voice     bool `json:"voice"`
		Interface bool `json:"interface"`
		Subtitles bool `json:"subtitles"`
	}

	GameLangs struct {
		EN Langs `json:"en"`
		RU Langs `json:"ru"`
		FR Langs `json:"fr"`
		ES Langs `json:"es"`
		DE Langs `json:"de"`
		IT Langs `json:"it"`
		PT Langs `json:"pt"`
	}

	GameReviews []GameReview
	GameReview  struct {
		PressName string `json:"press_name"`
		Link      string `json:"link"`
		Score     string `json:"score"`
		Quote     string `json:"quote"`
	}

	Socials struct {
		Facebook string
		Twitter  string
	}
)
