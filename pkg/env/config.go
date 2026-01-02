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
