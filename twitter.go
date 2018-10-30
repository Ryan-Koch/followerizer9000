package main

import (
	"fmt"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var client twitter.Client
var limitTracker LimitTracker

// LimitTracker -- used to make sure we don't over do the API requests and make Twitter unhappy.
type LimitTracker struct {
	Operations    int
	WindowStarted time.Time
}

func connectAndAuthorize() {
	config := oauth1.NewConfig(Settings.ConsumerKey, Settings.ConsumerSecret)
	token := oauth1.NewToken(Settings.AccessToken, Settings.AccessKey)

	httpClient := config.Client(oauth1.NoContext, token)

	client = *twitter.NewClient(httpClient)
}

func searchForUsers() []string {
	var users []string

	for _, val := range Settings.TweetInterests {
		search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
			Query: val,
		})

		if err != nil {
			fmt.Println("Problem searching Tweets", err)
		}

		for _, item := range search.Statuses {
			if item.User.FollowersCount >= Settings.FollowerThreshold {
				users = append(users, item.User.ScreenName)
			}
		}

	}

	return users
}
