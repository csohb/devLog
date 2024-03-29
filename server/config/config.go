package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Log struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
	Name  string `yaml:"name"`
}

type Database struct {
	DSN     string `yaml:"dsn"`
	MaxIdle int    `yaml:"max_idle"`
	MaxConn int    `yaml:"max_conn"`
}

type Kafka struct {
	brokers []string `yaml:"brokers"`
}

type Redis struct {
	Host string `yaml:"host"`
}

type Server struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	SSLUse bool   `yaml:"ssl_use"`
	SSLKey string `yaml:"ssl_key"`
	SSLCrt string `yaml:"ssl_crt"`
}

type Web struct {
	BaseURL           string `yaml:"base_url"`
	App               string `yaml:"app"`
	PublicPath        string `yaml:"public_path"`
	SessionTimeout    int    `yaml:"session_timeout"`
	DownloadDirectory string `yaml:"download_directory"`
}

type Email struct {
	YeongEmail string `yaml:"yeong_email"`
	YeongPwd   string `yaml:"yeong_pwd"`
	YujinEmail string `yaml:"yujin_email"`
	YujinPwd   string `yaml:"yujin_pwd"`
}

type S3 struct {
	Bucket      string `yaml:"bucket"`
	CfgFile     string `yaml:"cfg_file"`
	Credentials string `yaml:"credentials"`
}

type Config struct {
	ServerName string   `yaml:"server_name"`
	Log        Log      `yaml:"log"`
	Database   Database `yaml:"database"`
	Kafka      Kafka    `yaml:"kafka"`
	Redis      Redis    `yaml:"redis"`
	Server     Server   `yaml:"server"`
	Web        Web      `yaml:"web"`
	Email      Email    `yaml:"email"`
	S3         S3       `yaml:"s3"`
}

type CDNInfo struct {
	URL                string `yaml:"url"`
	User               string `yaml:"user"`
	Password           string `yaml:"password"`
	HLSUser            string `yaml:"hls_user"`
	HLSPassword        string `yaml:"hls_password"`
	BasePath           string `yaml:"base_path"`
	DownloadBaseURL    string `yaml:"download_base_url"`
	DownloadHLSBaseURL string `yaml:"download_hls_base_url"`
}

func LoadConfig(filename string) (*Config, error) {
	var cfg *Config
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		return nil, err
	} else if err = yaml.Unmarshal(yamlFile, &cfg); err != nil {
		return nil, err
	} else {
		return cfg, nil
	}
}
