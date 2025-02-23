package di

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/drivers/db"
	"github.com/famiphoto/famiphoto/api/drivers/storage"
	"github.com/valkey-io/valkey-go"
)

var mySQLCluster db.Cluster

func NewMySQLCluster() db.Cluster {
	if mySQLCluster != nil {
		return mySQLCluster
	}
	c := db.NewMySQLClient(
		config.Env.MySQLHostName,
		config.Env.MySQLPort,
		config.Env.MySQLUser,
		config.Env.MySQLPassword,
		config.Env.MySQLDatabase,
	)
	mySQLCluster = db.NewCluster("famiphoto_mysql", c)
	return mySQLCluster
}

var localStorage storage.Client

func NewLocalStorage() storage.Client {
	if localStorage != nil {
		return localStorage
	}

	localStorage = storage.NewLocalStorage()
	return localStorage
}

var esClient *elasticsearch.Client
var esTypedClient *elasticsearch.TypedClient

func NewElasticSearchClient() *elasticsearch.Client {
	if esClient != nil {
		return esClient
	}

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:              config.Env.ElasticsearchAddresses,
		Username:               config.Env.ElasticsearchUserName,
		Password:               config.Env.ElasticsearchPassword,
		CertificateFingerprint: config.Env.ElasticsearchFingerPrint,
	})
	if err != nil {
		panic(err)
	}

	esClient = client
	return esClient
}

func NewTypesElasticSearchClient() *elasticsearch.TypedClient {
	if esTypedClient != nil {
		return esTypedClient
	}

	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses:              config.Env.ElasticsearchAddresses,
		Username:               config.Env.ElasticsearchUserName,
		Password:               config.Env.ElasticsearchPassword,
		CertificateFingerprint: config.Env.ElasticsearchFingerPrint,
	})
	if err != nil {
		panic(err)
	}
	esTypedClient = client
	return esTypedClient
}

var sessionDB valkey.Client

func NewSessionDB() valkey.Client {
	if sessionDB != nil {
		return sessionDB
	}

	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: config.Env.SessionDBAddresses,
	})
	if err != nil {
		panic(err)
	}

	sessionDB = client
	return sessionDB
}
