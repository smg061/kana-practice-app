package models

type Classification string

func (c Classification) IsValid() bool {

	switch c {
	case "Katakana", "Hiragana", "Kanji":
		return true
	default:
		return false
	}

}

type Kana struct {
	Representation string         `json:"representation"`
	Romaji         string         `json:"romaji"`
	Classification Classification `json:"classification"`
}
