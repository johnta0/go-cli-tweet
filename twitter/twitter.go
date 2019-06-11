package twitter

import (
	"log"
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/ChimeraCoder/anaconda"
)

// Tweet sends tweet text to twitter.com via API keys
// and return error if it occured
func Tweet(c *cli.Context) error {
	if c.Args().Get(0) == "" || c.NArg() == 0 {
		log.Fatalf("Tweet content is blank.")
		return nil
	}

	tweet := c.Args().Get(0)
	err := sendTweet(tweet)	
	if err != nil {
		return err
	}
	return nil
}

func sendTweet(tweet string) error {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CLI_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CLI_CONSUMER_SECRET_KEY"))

	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_CLI_ACCESS_TOKEN"), os.Getenv("TWITTER_CLI_ACCESS_SECRET"))
	_, err := api.PostTweet(tweet, nil)
	if err != nil {
		return err
	}
	fmt.Println("Tweet successfully sent!")
	return nil
}
