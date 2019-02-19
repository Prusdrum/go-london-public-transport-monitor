package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type TubeDisruptionResponse struct {
	Category string `json:"category"`
}

type TubeStatusResponse struct {
	Type        string                   `json:"$type"`
	Id          string                   `json:"id"`
	Name        string                   `json:"name"`
	ModeName    string                   `json:"modeName"`
	Disruptions []TubeDisruptionResponse `json:"disruptions"`
}

func GetTubesStatus() []TubeStatusResponse {
	var apiKey = os.Getenv("API_KEY")
	var appId = os.Getenv("APP_ID")
	var baseUrl = "https://api.tfl.gov.uk"
	var queryString = fmt.Sprintf("?app_id=%s&app_key=%s", appId, apiKey)
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+"/Line/Mode/tube,bus/Status"+queryString, nil)

	req.Header.Add("content-type", "application/json")
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	body, readErr := ioutil.ReadAll(response.Body)

	if readErr != nil {
		fmt.Println(readErr)
	}
	tubes := make([]TubeStatusResponse, 0)
	json.Unmarshal(body, &tubes)

	return tubes
}
