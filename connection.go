package gosdk

import (
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
	"github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/opensearchtransport"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"net/http"
	"os"
	"strconv"
)

type OpensearchConfig struct {
	Host               string `mapstructure:"host"`
	Username           string `mapstructure:"username"`
	Password           string `mapstructure:"password"`
	InsecureSkipVerify bool   `mapstructure:"skip-ssl"`
}

func InitOpenSearchClient(conf *OpensearchConfig, isDebug bool) (*opensearch.Client, error) {
	opensearchConf := opensearch.Config{
		Addresses:     []string{conf.Host},
		Username:      conf.Username,
		Password:      conf.Password,
		RetryOnStatus: []int{502, 503, 504, 429},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: conf.InsecureSkipVerify},
		},
	}

	if isDebug {
		opensearchConf.Logger = &opensearchtransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		}
	}

	return opensearch.NewClient(opensearchConf)
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
		gormConf.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
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

type JaegerConfig struct {
	Host        string `mapstructure:"host"`
	Environment string `mapstructure:"env"`
	ServiceName string `mapstructure:"service-name"`
}

type CassandraConfig struct {
	Hosts    []string `mapstructure:"hosts"`
	User     string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`
	Keyspace string   `mapstructure:"keyspace"`
}

func InitCassandraConnection(conf *CassandraConfig) (*gocql.Session, error) {
	cluster := gocql.NewCluster(conf.Hosts...)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: conf.User,
		Password: conf.Password,
	}

	cluster.Keyspace = conf.Keyspace

	return cluster.CreateSession()
}
