package Models

import (
	"github.com/olivere/elastic"
)


func GetElasticCon(url string) *elastic.Client {
	client, err := elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL(url));
	if err != nil {
		panic(err);
	}
	return client
}