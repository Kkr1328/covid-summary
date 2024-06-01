package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "log"
)

const CovidDataURL = "https://static.wongnai.com/devinterview/covid-cases.json"

type Case struct {
    ConfirmedDate string  `json:"ConfirmedDate"`
    No            string  `json:"No"`
    Age           *int    `json:"Age"`
    Gender        string  `json:"Gender"`
    GenderEn      string  `json:"GenderEn"`
    Nation        string  `json:"Nation"`
    NationEn      string  `json:"NationEn"`
    Province      string  `json:"Province"`
    ProvinceId    int     `json:"ProvinceId"`
    District      string  `json:"District"`
}

type Summary struct {
    Province map[string]int `json:"province"`
    AgeGroup map[string]int `json:"ageGroup"`
}

func GetCovidSummary(context *gin.Context) {
    response, err := http.Get(CovidDataURL)
    if err != nil {
        log.Fatalf("Failed to fetch data: %v", err)
    }
    defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		context.JSON(response.StatusCode, gin.H{"error": "Failed to fetch data"})
		return
	}

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalf("Failed to read response body: %v", err)
    }

	var data struct {
		Data []Case `json:"data"`
	}
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Fatalf("Failed to unmarshal JSON: %v", err)
    }

    summary := Summary{
        Province: make(map[string]int),
        AgeGroup: make(map[string]int),
    }

    for _, content := range data.Data {
        summary.Province[content.Province]++

        ageGroup := "N/A"
        if content.Age != nil {
            age := *content.Age
            switch {
            case age <= 30:
                ageGroup = "0-30"
            case age <= 60:
                ageGroup = "31-60"
            default:
                ageGroup = "60+"
            }
        }
        summary.AgeGroup[ageGroup]++
    }

    context.JSON(http.StatusOK, summary)
}
