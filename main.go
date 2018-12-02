package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

type TwitterSecret struct {
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
}

func loadSecret(authInfo *TwitterSecret) {
	authInfo.ConsumerKey = os.Getenv("CONSUMER_KEY")
	authInfo.ConsumerSecret = os.Getenv("CONSUMER_SECRET")
	authInfo.AccessToken = os.Getenv("ACCESS_TOKEN")
	authInfo.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
}

func main() {
	var authInfo TwitterSecret
	anaconda.SetConsumerKey(authInfo.ConsumerKey)
	anaconda.SetConsumerKey(authInfo.ConsumerSecret)
	api := anaconda.NewTwitterApi(authInfo.AccessToken, authInfo.AccessTokenSecret)
	fmt.Printf("%+v", api)
}
