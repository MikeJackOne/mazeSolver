package main

import (
	"encoding/json"
	"log"
	"strings"
)

func main() {

	b := `{"forward":"tiger", "left":{"forward":{"upstairs":"exit"}, "left":"dragon"}, "right":{"forward":
		"dead end"}}`

	b = `{"forward":"tiger", "left": "ogre", "right":"demon"}`

	store := make(map[string]interface{})

	err := json.Unmarshal([]byte(b), &store)
	if err != nil {
		log.Fatal("error while converting json to map!", err)
	}

	log.Println(FindWay(store))
}

type Route struct {
	Route []string
	Rest  interface{}
}

// FindWay get map[string]interface{} input and return. see main_test for more usage
func FindWay(maze map[string]interface{}) string {
	result := findWay(maze)

	if len(result) == 0 {
		return "Sorry"
	} else {
		return `["` + strings.Join(result, `","`) + `"]`
	}
}

func findWay(maze map[string]interface{}) []string {

	var processList []Route

	for k, v := range maze {
		processList = append(processList, Route{[]string{k}, v})
	}

	for len(processList) > 0 {
		route := processList[0]
		if len(processList) == 1 {
			processList = []Route{}
		} else {
			processList = processList[1:]
		}
		switch route.Rest.(type) {
		case string:
			if route.Rest.(string) == "exit" {
				return route.Route
			}
		case map[string]interface{}:
			tmpMaze := route.Rest.(map[string]interface{})
			for k, v := range tmpMaze {
				processList = append(processList, Route{append(route.Route, k), v})
			}
		}
	}

	return []string{}

}
