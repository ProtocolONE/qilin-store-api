package dto

type LanguagesSupportDTO struct {
	Audio []string `json:"audio"`
	Text  []string `json:"text"`
}

type RequirementsDTO struct {
	OS        string `json:"os"`
	CPU       string `json:"cpu"`
	RAM       string `json:"ram"`
	GPU       string `json:"gpu"`
	DiskSpace string `json:"disk_space"`
}

type SystemsDTO struct {
	Minimal     RequirementsDTO `json:"minimal"`
	Recommended RequirementsDTO `json:"recommended"`
}

type GameRequirementsDTO struct {
	Languages LanguagesSupportDTO   `json:"languages"`
	Systems   map[string]SystemsDTO `json:"systems"`
}

type LinkDTO struct {
	Title string `json:"title"`
	Id    string `json:"id"`
}

type GameDTO struct {
	Name         string              `json:"name"`
	Preview      string              `json:"preview"`
	Title        string              `json:"title"`
	Rating       float32             `json:"rating"`
	Price        float32             `json:"price"`
	Description  string              `json:"description"`
	Platforms    []string            `json:"platforms"`
	Tags         []LinkDTO           `json:"tags"`
	ReleaseDate  string              `json:"releaseDate"`
	Developer    LinkDTO             `json:"developer"`
	Publisher    LinkDTO             `json:"publisher"`
	Requirements GameRequirementsDTO `json:"requirements"`
}
