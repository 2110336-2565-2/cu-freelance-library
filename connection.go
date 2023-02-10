package gosdk

import (
	"fmt"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
)

type ElasticsearchConfig interface {
	GetConfig() *elasticsearchConfig
}

type elasticsearchConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func (c *elasticsearchConfig) GetConfig() any {
	return c
}

func InitElasticDefaultClient(conf ElasticsearchConfig, isDebug bool) (*elasticsearch.Client, error) {
	config := conf.GetConfig()

	esConf := elasticsearch.Config{
		Addresses: []string{config.Host},
		Username:  config.Username,
		Password:  config.Password,
	}

	if isDebug {
		esConf.Logger = &elastictransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		}
	}

	client, _ := elasticsearch.NewClient(esConf)

	if _, err := client.Info(); err != nil {
		return nil, err
	}

	return client, nil
}

func InitElasticTypedClient(conf ElasticsearchConfig, isDebug bool) (*elasticsearch.TypedClient, error) {
	config := conf.GetConfig()

	esConf := elasticsearch.Config{
		Addresses: []string{config.Host},
		Username:  config.Username,
		Password:  config.Password,
	}

	if isDebug {
		esConf.Logger = &elastictransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		}
	}

	return elasticsearch.NewTypedClient(esConf)
}

type RabbitMQConfig interface {
	GetConfig() *rabbitMQConfig
}

type rabbitMQConfig struct {
	Host  string `mapstructure:"host"`
	VHost string `mapstructure:"vhost"`
}

func (c *rabbitMQConfig) GetConfig() any {
	return c
}

func InitRabbitMQConnection(conf RabbitMQConfig) (*amqp.Connection, error) {
	config := conf.GetConfig()

	return amqp.DialConfig(config.Host, amqp.Config{
		Vhost: config.VHost,
	})
}

type PostgresDatabaseConfig interface {
	GetConfig() *postgresDatabaseConfig
}

type postgresDatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSL      string `mapstructure:"ssl"`
}

func (c *postgresDatabaseConfig) GetConfig() any {
	return c
}

func InitPostgresDatabase(conf PostgresDatabaseConfig, isDebug bool) (db *gorm.DB, err error) {
	config := conf.GetConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, strconv.Itoa(config.Port), config.User, config.Password, config.Name, config.SSL)

	gormConf := &gorm.Config{}

	if !isDebug {
		gormConf.Logger = logger.Default.LogMode(logger.Silent)
	}

	db, err = gorm.Open(postgres.Open(dsn), gormConf)
	if err != nil {
		return nil, err
	}

	return
}

type RedisConfig interface {
	GetConfig() *redisConfig
}

type redisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func (c *redisConfig) GetConfig() any {
	return c
}

func InitRedisConnect(conf RedisConfig) (cache *redis.Client, err error) {
	config := conf.GetConfig()

	cache = redis.NewClient(&redis.Options{
		Addr: config.Host,
		DB:   config.DB,
	})

	return
}
