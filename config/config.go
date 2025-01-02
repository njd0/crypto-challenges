package config

import (
	"encoding/json"
	"os"
)

type Set1Chal1 struct {
	Input string	`json:"input"`
	Output string		`json:"output"`
}

type Set1Chal2 struct {
	Inputs []string	`json:"inputs"`
	Output string		`json:"output"`
}

type Set1Chal3 struct {
	Input string	`json:"input"`
}

type Set1Chal4 struct {
	Input string	`json:"input"`
}

type Set1Chal5Inputs struct {
	Text string	`json:"text"`
	Key string	`json:"key"`
}

type Set1Chal5 struct {
	Inputs Set1Chal5Inputs	`json:"inputs"`
	Output string						`json:"output"`
}

type Config struct {
	Set1Chal1 Set1Chal1 `json:"set1chal1"`
	Set1Chal2 Set1Chal2 `json:"set1chal2"`
	Set1Chal3 Set1Chal3 `json:"set1chal3"`
	Set1Chal4 Set1Chal4 `json:"set1chal4"`
	Set1Chal5 Set1Chal5 `json:"set1chal5"`
}

func GetConfig() (Config) {
	// Open the file
	file, err := os.Open("./config/config.json")
	if err != nil {
			panic(err)
	}
	defer file.Close()

	// Parse the JSON
	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
			panic(err)
	}

	return config
}