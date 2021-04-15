package twitter

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type tweeter struct {
	ApiKey       string
	ApiSecret    string
	AccessToken  string
	AccessSecret string
}

func New(apiKey string, apiSecret string, accessToken string, accessSecret string) *tweeter {
	return &tweeter{ApiKey: apiKey, ApiSecret: apiSecret, AccessToken: accessToken, AccessSecret: accessSecret}
}

func (t *tweeter) Tweet(message string) {
	var tweetMessage string

	config := oauth1.NewConfig(t.ApiKey, t.ApiSecret)
	token := oauth1.NewToken(t.AccessToken, t.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Get latest tweet
	tweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: "VaccineDk",
		Count:      1,
	})

	if err != nil {
		log.Fatal(err)
	}

	if len(tweets) == 0 {
		tweetMessage = message
	} else {
		if tweets[0].Text != message {
			tweetMessage = message
		}
	}

	// Send a Tweet
	if tweetMessage != "" {
		_, _, err := client.Statuses.Update(tweetMessage, nil)

		if err != nil {
			log.Fatal(err)
		}
	}
}
