package models

import (
	"database/sql"
)

type Classification int


type Kana struct {
	Representation string         `json:"representation"`
	Romaji         string         `json:"romaji"`
	Classification Classification `json:"classification"`
	Initial 	   string 		  `json:"initial"`
}


type KanaModel struct {
	DB *sql.DB
}

func (km *KanaModel) GetAllKana() ([]Kana, error) {

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
	return kanaResult, nil

}
	
func (km *KanaModel) GetKanaByClass(classification string) ([]Kana, error) {
	statement := "SELECT representation, romaji, classification, initial FROM kana WHERE classification = $1"
	rows, err := km.DB.Query(statement, classification)

	if err != nil {
		return nil, err
	}
	kanaResult, err := km.parseKanaRows(rows)
	if err != nil {
		return nil, err
	}
	return kanaResult, nil

}

func (km *KanaModel) GetKanagetKanaByInitial (initial string) ([]Kana, error) {
	statement := "SELECT representation, romaji, classification, initial FROM kana WHERE initial = $1"
	rows, err := km.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	kanaResult, err := km.parseKanaRows(rows)
	if err != nil {
		return nil, err
	}
	return kanaResult, nil
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