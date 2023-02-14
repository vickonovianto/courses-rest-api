package config

import (
	"courses-api/config/postgres"
	"os"
	"strconv"

	"gorm.io/gorm"
)

type (
	config struct {
	}

	Config interface {
		ServicePort() int
		Database() *gorm.DB
	}
)

func NewConfig() Config {
	return &config{}
}

func (c *config) Database() *gorm.DB {
	return postgres.InitGorm()
}

func (c *config) ServicePort() int {
	v := os.Getenv("PORT")
	port, _ := strconv.Atoi(v)

	return port
}
