package inlogrus

import (
	"github.com/sirupsen/logrus"
	"github.com/sohlich/elogrus"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

// send log info to elasticsearch

type Tweet struct {
	User     string
	Message  string
	Retweets int
}

type msg struct {
	Host      string
	Timestamp string `json:"@timestamp"`
	Message   string
	Data      logrus.Fields
	Level     string
}

func Elastic() *elogrus.ElasticHook {
	client, err := elastic.NewClient(elastic.SetURL())
	if err != nil {
		log.Panic(err)
	}
	// Index a tweet (using JSON serialization)
	//tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	elkHook, err := elogrus.NewElasticHook(client, "localhost", logrus.DebugLevel, "mylog")
	return elkHook
}
