package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
)

type Config struct {
	DBUser     string `json:"db_user" flag:"db-user" env:"DB_USER"`
	DBPassword string `json:"db_password" flag:"db-password" env:"DB_PASSWORD"`
	DBName     string `json:"db_name" flag:"db-name" env:"DB_NAME"`
	DBHost     string `json:"db_host" flag:"db-host" env:"DB_HOST"`
	DBPort     string `json:"db_port" flag:"db-port" env:"DB_PORT"`
	HTTPPort   string `json:"http_port" flag:"http-port" env:"HTTP_PORT"`
}

func (c *Config) setDefaults() {
	c.DBUser = "user"
	c.DBPassword = "password"
	c.DBName = "todo_db"
	c.DBHost = "localhost"
	c.DBPort = "5432"
	c.HTTPPort = "8080"
}

func Load() (Config, error) {
	var cfg Config
	cfg.setDefaults()

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path to JSON config")
	flag.StringVar(&configPath, "c", "", "Short alias for --config")

	t := reflect.TypeOf(cfg)
	v := reflect.ValueOf(&cfg).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		flagTag := field.Tag.Get("flag")
		if flagTag == "" {
			continue
		}
		flag.StringVar(v.Field(i).Addr().Interface().(*string), flagTag, v.Field(i).String(), "")
	}

	flag.Parse()

	if configPath == "" {
		configPath = os.Getenv("CONFIG")
	}

	if configPath != "" {
		jsonCfg, err := loadFromJSON(configPath)
		if err != nil {
			return Config{}, fmt.Errorf("failed to load JSON config: %w", err)
		}
		mergeConfigs(&cfg, &jsonCfg)
	}

	if err := loadFromEnv(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func loadFromJSON(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return cfg, nil
}

func loadFromEnv(cfg *Config) error {
	t := reflect.TypeOf(*cfg)
	v := reflect.ValueOf(cfg).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		envTag := field.Tag.Get("env")
		if envTag == "" {
			continue
		}
		if val := os.Getenv(envTag); val != "" {
			v.Field(i).SetString(val)
		}
	}
	return nil
}

func mergeConfigs(target, from *Config) {
	t := reflect.TypeOf(*target)
	vTarget := reflect.ValueOf(target).Elem()
	vFrom := reflect.ValueOf(from).Elem()

	for i := 0; i < t.NumField(); i++ {
		val := vTarget.Field(i)
		if val.String() == "" {
			val.SetString(vFrom.Field(i).String())
		}
	}
}

func (c Config) DBURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

func (c Config) HTTPAddr() string {
	if c.HTTPPort == "" {
		return ":8080"
	}
	return ":" + c.HTTPPort
}
