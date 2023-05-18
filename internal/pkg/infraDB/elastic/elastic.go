package elastic

import (
	"fmt"
	"github.com/Hamster601/fastweb/configs"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

var ElasticCLient elasticsearch8.Client

func New() {
	hostAddr := configs.Get().Elastic.Host
	elasticCfg := elasticsearch8.Config{
		Addresses:    []string{hostAddr},
		Username:     "",
		Password:     "",
		APIKey:       "",
		ServiceToken: "",
		CACert:       nil,
	}
	client, err := elasticsearch8.NewClient(elasticCfg)
	if err != nil {
		return
	}
	fmt.Println(client.Info)
}

func Close() {
}
