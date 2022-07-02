package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type KanaRecord struct {
	Representation string
	Romaji         string
	Classification int
	Initial 	   string
}

func ReadKanaRecords() ([]KanaRecord, error) {

	filename := "kana-data.csv"

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	_, err = csvReader.Read()
	
	if err != nil {
		log.Fatal(err)
	}
	reads, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	inserts := []KanaRecord{}
	for _, row := range reads {
		kana, err := newKanaRecord(row)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		inserts = append(inserts, *kana)
		fmt.Printf("%v", *kana)
	}

	return inserts, nil

}

func newKanaRecord(row []string) (*KanaRecord, error) {
	var kana KanaRecord
	kana.Representation = row[0]
	kana.Romaji = row[1]
	c, err := strconv.Atoi(row[2])
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	kana.Classification = c
	kana.Initial = row[3]

	return &kana, nil

}

