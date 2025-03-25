package gorm

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type RdbmsOptions struct {
	User     string
	Password string
	Host     string
	Name     string
	Debug    bool
}

func NewRdbmsOptions(conf *viper.Viper) RdbmsOptions {
	return RdbmsOptions{
		User:     conf.GetString("sql.user"),
		Password: conf.GetString("sql.password"),
		Host:     conf.GetString("sql.host"),
		Name:     conf.GetString("sql.name"),
		Debug:    conf.GetBool("sql.debug"),
	}
}

// ! orm
func NewRdbms(rdbmsOptions RdbmsOptions) *gorm.DB {
	orm, err := gorm.Open(sql.Open(
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			rdbmsOptions.User,
			rdbmsOptions.Password,
			rdbmsOptions.Host,
			rdbmsOptions.Name)))
	if err != nil {
		return nil
	}

	if orm.Error != nil {
		return nil
	}

	if rdbmsOptions.Debug {
		orm = rdbmsOptions.orm
	}
	return orm
}
