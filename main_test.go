package main

import (
	"encoding/json"
	"testing"
)

func TestFindWaySorry(t *testing.T) {
	b := `{"forward":"tiger", "left": "ogre", "right":"demon"}`

	store := make(map[string]interface{})

	json.Unmarshal([]byte(b), &store)
	result := FindWay(store)
	if result != "Sorry" {
		t.Fatalf(`failed sorry test,get %s expected Sorry:`, result)
	}
}

func TestFindWayNullTest(t *testing.T) {
	store := make(map[string]interface{})

	result := FindWay(store)
	if result != "Sorry" {
		t.Fatalf(`failed null test,get %s expected Sorry:`, result)
	}
}

func TestFindWaySuccessful(t *testing.T) {
	b := `{"forward":"tiger", "left":{"forward":{"upstairs":"exit"}, "left":"dragon"}, "right":{"forward":
		"dead end"}}`

	store := make(map[string]interface{})

	json.Unmarshal([]byte(b), &store)
	result := FindWay(store)
	if result != `["left","forward","upstairs"]` {
		t.Fatalf(`failed success test,get %s expected ["left","forward","upstairs"]:`, result)
	}
}
