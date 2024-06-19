package helpers

import (
	"fmt"
	"net/url"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//model "issue-tracker.com/models"
)

type Repository interface {
	Get() *gorm.DB
}

type repository struct {
	DB *gorm.DB
}

func ConnectToDb() Repository {
	var err error

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := url.URL{
		User:     url.UserPassword(user, password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s,%s", host, port),
		Path:     database,
		RawQuery: (&url.Values{"sslmode": []string{"require"}}).Encode(),
	}

	//"host user=avnadmin password= dbname= port=21609 sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	//db, err := gorm.Open("postgres", dsn.String())

	if err != nil {
		panic("Failed to connect to a db")
	}

	return &repository{DB: db}

}

func (d *repository) Get() *gorm.DB {

	return d.DB
}
