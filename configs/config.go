package configs

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"paint/pkg/repo"
)

type Config struct {
	Server struct {
		Host           string `yaml:"host"`
		Port           int    `yaml:"port"`
		MetricAddrPort int    `yaml:"metricAddrPort"`
		Timeout        struct {
			Server int `yaml:"server"`
			Write  int `yaml:"write"`
			Read   int `yaml:"read"`
			Idle   int `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`

	DataBase struct {
		Host          string `yaml:"host"`
		Port          int    `yaml:"port"`
		User          string `yaml:"user"`
		Pass          string `yaml:"pass"`
		Name          string `yaml:"name"`
		MaxIdleConns  int    `yaml:"maxIdleConns"`
		MaxOpenConns  int    `yaml:"maxOpenConns"`
		MigrationPath string `yaml:"migrationPath"`
	} `yaml:"database"`

	ImageProcServer struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		LogLevel string `yaml:"logLevel"`
	} `yaml:"imageProc"`
}

func Init() (*Config, error) {
	cfgPath, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := newConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	return cfg, err
}

func (c *Config) GetDbConfig() *repo.Config {
	return &repo.Config{
		Host:          c.DataBase.Host,
		Port:          c.DataBase.Port,
		Name:          c.DataBase.Name,
		User:          c.DataBase.User,
		Pass:          c.DataBase.Pass,
		MaxIdleConns:  c.DataBase.MaxOpenConns,
		MaxOpenConns:  c.DataBase.MaxOpenConns,
		MigrationPath: c.DataBase.MigrationPath,
	}
}

func newConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func parseFlags() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "./configs/config.yml", "path to config file")

	flag.Parse()

	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}
