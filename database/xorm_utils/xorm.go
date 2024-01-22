package xorm_utils

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
)

type XormConfig struct {
	User     string `yaml:"user" mapstructure:"user"`
	Password string `yaml:"password" mapstructure:"password"`
	Host     string `yaml:"host" mapstructure:"host"`
	Port     string `yaml:"port" mapstructure:"port"`
	Database string `yaml:"database" mapstructure:"database"`
	ShowSql  bool   `yaml:"show_sql" mapstructure:"show_sql"`
}

func NewXorm(config XormConfig) *xorm.Engine {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
	)

	log.Printf("database connection string is: %s", psqlInfo)
	x, err := xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		log.Panicf("连接数据库失败：%s", err)
	}
	if config.ShowSql {
		x.ShowSQL(true)
	}
	return x
}
