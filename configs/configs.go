package configs

import (
	"fmt"
	"kayky18/api-banking/internal/entity"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Conf struct {
	JwtSecret   string `mapstructure:"JwtSecretKey"`
	JwtExpireIn int    `mapstructure:"JwtExpireIn"`
	TokenAuth   *jwt.Token
}

func InitDataBase() (*gorm.DB, error) {
	dbPath := "db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		fmt.Println("Database file not found, creating..")

		//Create the database path
		_, err := os.Stat("./db")
		if os.IsNotExist(err) {
			err = os.Mkdir("./db", os.ModePerm)
			if err != nil {
				err = fmt.Errorf("ERRO CREATING DATABASE PATH: %v", err.Error())
				return nil, err
			}
		}

		//Creating the database file
		file, err := os.Create(dbPath)
		if err != nil {
			err = fmt.Errorf("ERRO CREATING DATABASE FILE: %v", err.Error())
			return nil, err
		}

		//Close the file
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Transaction{})

	return db, nil

}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf

	viper.SetConfigFile("app_config")

	viper.AddConfigPath(path)

	viper.SetConfigType("env")

	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	cfg.TokenAuth = jwt.New(jwt.SigningMethodHS256)
	return cfg, nil
}
