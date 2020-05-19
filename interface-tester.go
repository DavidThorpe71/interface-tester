package main

import (
	"encoding/json"
	"fmt"
)

var testString = `
	{
		"title": "newTestTitle",
		"meta": {
			"id": 123,
			"urn": "1",
			"date": "19/12/2020",
			"entity": "article"
		},
		"text": "new test text"
	}
`

var testString2 = `
	{
		"collectionTitle": "newTestCollectionTitle",
		"meta": {
			"id": 456,
			"urn": "2",
			"date": "19/12/2020",
			"entity": "collection"
		},
		"collectionText": "new test collection text"
	}
`

type GenericFunctions interface {
	getEnt(string) (string, error)
}

type Meta struct {
	ID     int    `json:"id"`
	Date   string `json:"date"`
	Urn    string `json:"urn"`
	Entity string `json:"entity"`
}

type Article struct {
	Title string `json:"title"`
	Meta  Meta   `json:"meta"`
	Text  string `json:"text"`
}
type Collection struct {
	CollectionTitle string `json:"collectionTitle"`
	Meta            Meta   `json:"meta"`
	CollectionText  string `json:"collectionText"`
}

type MetaOnly struct {
	Meta Meta `json:"meta"`
}

func getEntity(t GenericFunctions, stt string) (string, error) {
	return t.getEnt(stt)
}

func (m MetaOnly) getEnt(stringToTest string) (string, error) {
	err := json.Unmarshal([]byte(stringToTest), &m)
	if err != nil {
		return "", err
	}
	return m.Meta.Entity, nil
}

func (a Article) getEnt(stringToTest string) (string, error) {
	return "hello", nil
}

func main() {

	var metaToTest MetaOnly
	var articleToTest Article

	stringsArray := []string{testString, testString2}

	for _, item := range stringsArray {
		entity, entityErr := getEntity(metaToTest, item)
		newEntity, newEntErr := getEntity(articleToTest, item)

		if entityErr != nil || newEntErr != nil {
			fmt.Println("Error: ", entityErr)
		}
		fmt.Println("David nbew test:", newEntity)

		if entity == "article" {
			var testArticle Article
			err := json.Unmarshal([]byte(item), &testArticle)

			if err != nil {
				fmt.Println("Error: ", err)
			}

			fmt.Println("testArticle", testArticle)
		}

		if entity == "collection" {
			var testCollection Collection
			err := json.Unmarshal([]byte(item), &testCollection)

			if err != nil {
				fmt.Println("Error: ", err)
			}

			fmt.Println("testCollection", testCollection)
		}
	}

}
