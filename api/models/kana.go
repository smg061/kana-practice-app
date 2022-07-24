package models

import (
	"database/sql"
)


// this is the db entity
type Kana struct {
	Representation string         
	Romaji         string         
	Classification int 
	Initial 	   string 		  
}

type KanaJsonView struct {
	DisplayValue string         `json:"displayValue"`
	CorrectAnswer string        `json:"correctAnswer"`
	Classification string 		`json:"classification"`
	Initial 	   string 		`json:"initial"`
}

func (kn Kana) ToJsonView() (view *KanaJsonView) {
	jsonView  :=&KanaJsonView{
		DisplayValue: kn.Representation,
		CorrectAnswer: kn.Romaji,
		Initial: kn.Initial,
	}
	if kn.Classification == 1 {
		jsonView.Classification = "hiragana"
	} else {
		jsonView.Classification = "katakana"
	}
	return jsonView
}

type KanaModel struct {
	DB *sql.DB
}

func (km *KanaModel) GetAllKana() ([]KanaJsonView, error) {

	statement := "SELECT representation, romaji, classification, initial FROM kana"
	rows, err := km.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	kanaResult, err := km.parseKanaRows(rows)
	if err != nil {
		return nil, err
	}
	var kanaJsonViews []KanaJsonView

	for _, k := range kanaResult {
		kanaJsonViews = append(kanaJsonViews, *k.ToJsonView())
	}
	return kanaJsonViews, nil

}
	
func (km *KanaModel) GetByClass(classification float64) ([]KanaJsonView, error) {
	
	statement := "SELECT representation, romaji, classification, initial FROM kana WHERE classification = $1"
	rows, err := km.DB.Query(statement, classification)

	if err != nil {
		return nil, err
	}
	kanaResult, err := km.parseKanaRows(rows)
	var kanaJsonViews []KanaJsonView

	for _, k := range kanaResult {
		kanaJsonViews = append(kanaJsonViews, *k.ToJsonView())
	}
	if err != nil {
		return nil, err
	}
	
	return kanaJsonViews, nil

}

func (km *KanaModel) GetByInitial (initial string) ([]KanaJsonView, error) {
	statement := "SELECT representation, romaji, classification, initial FROM kana WHERE initial = $1"
	rows, err := km.DB.Query(statement, initial)
	if err != nil {
		return nil, err
	}
	kanaResult, err := km.parseKanaRows(rows)
	if err != nil {
		return nil, err
	}
	var kanaJsonViews []KanaJsonView

	for _, k := range kanaResult {
		kanaJsonViews = append(kanaJsonViews, *k.ToJsonView())
	}
	return kanaJsonViews, nil
}

func (km *KanaModel) parseKanaRows(input *sql.Rows) ([]Kana, error) {
	kanaResult := []Kana{}
	for input.Next() {
		var kana Kana
		err := input.Scan(&kana.Representation, &kana.Romaji, &kana.Classification, &kana.Initial)
		if err != nil {
			return nil, err
		}
		kanaResult = append(kanaResult, kana)
	}
	return kanaResult, nil
}