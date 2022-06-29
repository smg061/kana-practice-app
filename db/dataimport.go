package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	Hiragana = iota
	Katakana
	Kanji
)

type KanaRecord struct {
	Representation string
	Romaji         string
	Classification int
}

func main() {

	if len(os.Args) == 1 {
		log.Fatal("Please provide a filename")
	}
	filename := os.Args[1]

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	_, err = csvReader.Read()

	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range data {
		kana, err := newKanaRecord(row)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v", *kana)
	}

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

	return &kana, nil

}
