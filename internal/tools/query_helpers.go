package tools

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Aisle struct {
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

type Aisles struct {
	Aisles []Aisle `json:"aisles"`
}

func JsonMapToAisles(file_name string) (Aisles, map[string]int) {
	jsonFile, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var aisles Aisles
	aisleMap := make(map[string]int)

	json.Unmarshal(byteValue, &aisles)

	for index, value := range aisles.Aisles {
		aisleMap[value.Name] = index
	}

	return aisles, aisleMap
}

func QueryBuilder(aisles Aisles, aisleMap map[string]int, path []string, list string) string {
	query := "Sort the following grocery list by order of "

	for index, value := range path {
		if index != 0 {
			query += " then "
		}
		query += strings.Join(aisles.Aisles[aisleMap[value]].Items, ", ")
	}

	query += " : "
	query += list
	query += ". Just give a list do not show the categories."

	return query
}
