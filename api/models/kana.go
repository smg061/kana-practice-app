package models

import (
	"database/sql"
)




type Kana struct {
	Representation string         `json:"representation"`
	Romaji         string         `json:"romaji"`
	Classification int `json:"classification"`
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
	
func (km *KanaModel) GetByClass(classification int) ([]Kana, error) {
	
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

func (km *KanaModel) GetByInitial (initial string) ([]Kana, error) {
	statement := "SELECT representation, romaji, classification, initial FROM kana WHERE initial = $1"
	rows, err := km.DB.Query(statement, initial)
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