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

type ElasticsearchConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func InitElasticDefaultClient(conf *ElasticsearchConfig, isDebug bool) (*elasticsearch.Client, error) {
	esConf := elasticsearch.Config{
		Addresses: []string{conf.Host},
		Username:  conf.Username,
		Password:  conf.Password,
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

func InitElasticTypedClient(conf *ElasticsearchConfig, isDebug bool) (*elasticsearch.TypedClient, error) {
	esConf := elasticsearch.Config{
		Addresses: []string{conf.Host},
		Username:  conf.Username,
		Password:  conf.Password,
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

type RabbitMQConfig struct {
	Host  string `mapstructure:"host"`
	VHost string `mapstructure:"vhost"`
}

func InitRabbitMQConnection(conf *RabbitMQConfig) (*amqp.Connection, error) {
	return amqp.DialConfig(conf.Host, amqp.Config{
		Vhost: conf.VHost,
	})
}

type PostgresDatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSL      string `mapstructure:"ssl"`
}

func InitPostgresDatabase(conf *PostgresDatabaseConfig, isDebug bool) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", conf.Host, strconv.Itoa(conf.Port), conf.User, conf.Password, conf.Name, conf.SSL)

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

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func InitRedisConnect(conf *RedisConfig) (cache *redis.Client, err error) {
	cache = redis.NewClient(&redis.Options{
		Addr: conf.Host,
		DB:   conf.DB,
	})

	return
}
