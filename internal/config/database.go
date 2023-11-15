package config

import (
	"fmt"
	"os"

	"github.com/opengo-org/easy-netdisk/module"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dns string

func GetDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dns))
	if err != nil {
		fmt.Printf("Fail to connect database")
	}
	return db
}

func loadConfiguration() module.Configuration {
	var config module.Configuration

	work_dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Fail to get current work directory")
	}

	// Viper Read Configurations
	v := viper.New()
	v.SetConfigFile("config/config.yaml")
	v.AddConfigPath(work_dir)
	v.ReadInConfig()
	v.Unmarshal(&config)
	return config
}

// Init dns for database connection
func init() {
	cfg := loadConfiguration()
	dns = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		cfg.Database.Host, cfg.Database.User,
		cfg.Database.Password, cfg.Database.Dbname,
		cfg.Database.Port)
}
