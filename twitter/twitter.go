package twitter

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func Tweet(message string) {
	var tweetMessage string

	config := oauth1.NewConfig(os.Getenv("TWITTER_API_KEY"), os.Getenv("TWITTER_API_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
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

	// fmt.Println(tweets[0].Text)

	// Send a Tweet
	if tweetMessage != "" {
		_, _, err := client.Statuses.Update(tweetMessage, nil)

		if err != nil {
			log.Fatal(err)
		}
	}
}
