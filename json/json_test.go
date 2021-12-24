package main_test

import (
	"encoding/json"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	bytes := []byte("{\"Name\":\"Luca\",\"Lang\":[\"Java\",\"Python\",\"Go\"]}")
	var data map[string]interface{}
	_ = json.Unmarshal(bytes, &data)
	if data["Name"] != "Luca" {
		t.Errorf("expect: %v, actual: %v", "Luca", data["Name"])
	}
	lang := data["Lang"].([]interface{})
	if len(lang) != 3 {
		t.Errorf("expect: %v, actual: %v", 3, len(lang))
	}
	l := lang[2].(string)
	if l != "Go" {
		t.Errorf("expect: %v, actual: %v", "Go", l)
	}
}

func TestMarshal(t *testing.T) {
	type Developer struct {
		// 大写字母开头(public)的字段才能反序列化成功
		Name string
		Lang []string
	}
	dev := &Developer{Name: "Luca", Lang: []string{
		"Java", "Python", "Go",
	}}
	bytes, err := json.Marshal(dev)
	if err != nil {
		t.Error(err)
	}

	var data Developer
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		t.Error(err)
	}
	if data.Name != "Luca" {
		t.Errorf("expect: %v, actual: %v", "Luca", data.Name)
	}
	if data.Lang[2] != "Go" {
		t.Errorf("expect: %v, actual: %v", "Go", data.Lang[2])
	}
}
