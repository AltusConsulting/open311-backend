package connectors

import (
	"github.com/spf13/viper"
	elastic "gopkg.in/olivere/elastic.v5"
)

func Create311Client() (*elastic.Client, error) {
	client, err := elastic.NewSimpleClient(elastic.SetURL(viper.GetString("311.host")))
	return client, err
}
