package config

import "fmt"

type Config struct {
	DB     `yaml:"db"`
	Server `yaml:"server"`
}
type DB struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Server struct {
	Prot int `yaml:"port"`
}

func (db *DB) CreateDNS() string {
	dns := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, db.Username, db.Password, db.Host, db.Port, db.Database)
	return dns
}
