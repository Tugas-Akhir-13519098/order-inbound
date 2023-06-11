package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	RESTHost string `envconfig:"REST_HOST" default:"localhost"`
	RESTPort int    `envconfig:"REST_PORT" default:"8000"`

	KafkaHost       string `envconfig:"KAFKA_HOST" default:"localhost"`
	KafkaPort       string `envconfig:"KAFKA_PORT" default:"9092"`
	KafkaOrderTopic string `envconfig:"KAFKA_PRODUCT_TOPIC" default:"order"`

	ShopeeURL string `envconfig:"SHOPEE_URL" default:"https://fb724b04-b7cc-47d0-bba0-21e93dda67b8.mock.pstmn.io/api/v2/order/"`
}

// Get to get defined configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	return cfg
}
