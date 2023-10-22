package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

func main() {
	f, err := os.Open("data/data.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	objects := []*Object{}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '\t'
		return r
	})

	if err := gocsv.UnmarshalFile(f, &objects); err != nil {
		panic(err)
	}

	validContexts := make(map[string]map[string]interface{})
	for _, object := range objects {
		isValid := object.VerifyContextMap()

		if isValid {
			contextMap := object.ContextMap()
			validContexts[object.CODE1] = contextMap
		}
	}

	i := buildLatticeMinerInput(validContexts)

	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	fileName := settings()["name"].(string) + ".json"
	os.WriteFile(fileName, b, 0644)
}

func buildLatticeMinerInput(validContexts map[string]map[string]interface{}) *LatticeMinerInput {
	settings := settings()
	return &LatticeMinerInput{
		Name: settings["name"].(string),
		Objects: buildObjects(validContexts),
		Attributes: buildAttributes(validContexts, settings),
		Conditions: buildConditions(validContexts, settings),
		Relations: buildRelations(validContexts),
	}
}

func buildRelations(validContexts map[string]map[string]interface{}) [][][]string{
	relations := [][][]string{}
	for _, context := range validContexts {
		relationObject := [][]string{}
		relationWaves := []string{}
		wave := 0
		for _, value := range context {
			relation := getRelation(value.(string), wave)
			if relation != "" {
				relationWaves = append(relationWaves, relation)
			}

			if wave == 2 {
				relationObject = append(relationObject, relationWaves)
				wave = 0
				relationWaves = []string{}
			} else {
				wave = wave + 1
			}
		}
		relations = append(relations, relationObject)
	}

	return relations
}

func getRelation(value string, wave int) string {
	r := ""

	if value == "True" {
		switch wave {
		case 0:
			r = "t0"
		case 1:
			r = "t1"
		case 2:
			r = "t2"
		}
	} 

	return r
}

func buildConditions(validContexts map[string]map[string]interface{}, settings map[string]interface{}) []string{
	return settings["waves"].([]string)
}

func buildAttributes(validContexts map[string]map[string]interface{}, settings map[string]interface{}) []string{
	return settings["attrs"].([]string)
}

func buildObjects(validContexts map[string]map[string]interface{}) []string {
	var objects []string
	for k := range validContexts {
		objects = append(objects, k)
	}
	return objects
}

func settings() map[string]interface{} {
	return map[string]interface{}{
		"name": "fdtable",
		"noWaves": 3,
		"noAttr": 19,
		"attrs": []string{"HADS1", "HADS2", "HADS3", "HADS4", "HADS6", "HADS9", "HADS11", "HADS12", "HADS14", "SOMS1", "SOMS2", "SOMS3", "SOMS4", "SOMS8", "SOMS13", "SOMS14", "SOMS29", "SOMS30", "SOMS54"},
		"waves": []string{"t0", "t1", "t2"},
	}
}