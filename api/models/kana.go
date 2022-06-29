package models

type Classification int


type Kana struct {
	Representation string         `json:"representation"`
	Romaji         string         `json:"romaji"`
	Classification Classification `json:"classification"`
}
