package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/genai"
	"encoding/json"	
	"os"
	"io/ioutil"
	"strings"
)

type Aisle struct {
	Name string `json:"name"`
	Items []string `json:"items"`
}


type Aisles struct{
	Aisles []Aisle `json:"aisles"`
}

func main() {
	list := "bananas, oranges, salad, frozen pizza, pizza rolls, ice cream, butter, cottage cheese, salt, garlic powder, bread, eggs, soy sauce,  marinara"
	
	jsonFile, err := os.Open("aisles.json")
	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var aisles Aisles
	aisleMap := make(map[string]int)

	json.Unmarshal(byteValue, &aisles)

	for index, value := range aisles.Aisles {
		aisleMap[value.Name] = index
	}	

	var path []string
	path = append(path, "produce")
	path = append(path, "two_front")
	path = append(path, "three_front")
	path = append(path, "four_front")
	path = append(path, "two_back")
	path = append(path, "three_back")
	path = append(path, "four_back")
	path = append(path, "meatcase")
	path = append(path, "refridgerated")
	path = append(path, "frozen")
	path = append(path, "nine_front")
	path = append(path, "eight_front")
	path = append(path, "seven_front")
	path = append(path, "six_front")
	path = append(path, "five_front")
	path = append(path, "five_back")
	path = append(path, "six_back")
	path = append(path, "seven_back")
	path = append(path, "eight_back")
	path = append(path, "nine_back")

	query := "Please sort this the following grocery list by order of "
	
	for index, value := range path {
		if(index != 0){
			query += " then "
		}
		query += strings.Join(aisles.Aisles[aisleMap[value]].Items, ", ")
	}

	query += " : " 
	query += list
	query += ". Just give a list do not show the categories."

	geminiCall(query)
}

func geminiCall(query string) {
    ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text(query),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(result.Text())
}
