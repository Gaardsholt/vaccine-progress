package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Gaardsholt/vaccine-progress/twitter"
)

func main() {
	res, err := http.Get("https://static.editorialdev.tv2a.dk/assets/2020/covid19/vaccinesDashboard.json")
	if err != nil {
		log.Fatal("Unable to get data")
	}

	var result Vaccines
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Fatalf("Unable to decode data: %v", err)
		return
	}

	// Thank you for providing me with correct data types...
	fix_number := strings.Replace(result.Danmark.Vaccinedoneper100, ",", ".", -1)
	number, err := strconv.ParseFloat(fix_number, 64)

	if err != nil {
		log.Fatalf("Unable to parse number to float: %v", err)
	}

	progress := createProgressBar(number)

	fmt.Println(progress)

	apiKey := os.Getenv("TWITTER_API_KEY")
	apiSecret := os.Getenv("TWITTER_API_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	twitterClient := twitter.New(apiKey, apiSecret, accessToken, accessSecret)
	twitterClient.Tweet(progress)
}

func createProgressBar(percentage float64) string {
	var result string

	solids := int(math.Floor(percentage * 15 / 100))
	blanks := 15 - solids

	result += strings.Repeat("▓", solids)
	result += strings.Repeat("░", blanks)

	return fmt.Sprintf("%s %.2f%%", result, percentage)
}

type Vaccines struct {
	Danmark struct {
		Vaccinestarted       string `json:"vaccinestarted"`
		Vaccinestartedper100 string `json:"vaccinestartedper100"`
		Vaccinedone          string `json:"vaccinedone"`
		Vaccinedoneper100    string `json:"vaccinedoneper100"`
		Dosestotal           string `json:"dosestotal"`
		Dosestotalper100     string `json:"dosestotalper100"`
		Dosesdaily           string `json:"dosesdaily"`
		Updated              string `json:"updated"`
	} `json:"Danmark"`
	World struct {
		Vaccinestarted        string `json:"vaccinestarted"`
		Vaccinestartedper100  string `json:"vaccinestartedper100"`
		Fullyvaccinated       string `json:"fullyvaccinated"`
		Fullyvaccinatedper100 string `json:"fullyvaccinatedper100"`
		Updated               string `json:"updated"`
	} `json:"World"`
	Disclaimer string `json:"disclaimer"`
}
