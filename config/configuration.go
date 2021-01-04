package config

type Configuration struct {
	Bot   BotConfiguration
	Mongo MongoConfiguration
	Owner OwnerConfiguration
	AVWX  AVWXConfiguration
}