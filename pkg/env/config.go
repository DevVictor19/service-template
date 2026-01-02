package env

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server *Server
	Logger *Logger
	MySQL  *MySQL
}

type Server struct {
	Mode           string
	Addr           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
	CtxTimeout     time.Duration
}

type Logger struct {
	Level    string
	Encoding string
}

type MySQL struct {
	Host            string
	Port            string
	User            string
	DBname          string
	Password        string
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
}

func NewConfig() *Config {
	{
		return &Config{
			Server: &Server{
				Mode:           getString("SERVER_MODE"),
				Addr:           getString("SERVER_ADDR"),
				ReadTimeout:    time.Duration(getInt("SERVER_READ_TIMEOUT")) * time.Second,
				WriteTimeout:   time.Duration(getInt("SERVER_WRITE_TIMEOUT")) * time.Second,
				MaxHeaderBytes: getInt("SERVER_MAX_HEADER_BYTES"),
				CtxTimeout:     time.Duration(getInt("SERVER_CTX_TIMEOUT")) * time.Second,
			},
			Logger: &Logger{
				Level:    getString("LOGGER_LEVEL"),
				Encoding: getString("LOGGER_ENCODING"),
			},
			MySQL: &MySQL{
				Host:            getString("MYSQL_HOST"),
				Port:            getString("MYSQL_PORT"),
				User:            getString("MYSQL_USER"),
				DBname:          getString("MYSQL_DBNAME"),
				Password:        getString("MYSQL_PASSWORD"),
				MaxOpenConns:    getInt("MYSQL_MAX_OPEN_CONNS"),
				ConnMaxLifetime: time.Duration(getInt("MYSQL_CONN_MAX_LIFETIME")) * time.Second,
				MaxIdleConns:    getInt("MYSQL_MAX_IDLE_CONNS"),
				ConnMaxIdleTime: time.Duration(getInt("MYSQL_CONN_MAX_IDLE_TIME")) * time.Second,
			},
		}
	}
}

func getString(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("missing %s on .env file\n", key)
	}

	return val
}

func getInt(key string) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("missing %s on .env file\n", key)
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}

	return valAsInt
}
