package dbs

import (
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/martenwallewein/go-sample-microservice/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func GetDB() *gorm.DB {
	return dbInstance
}

func setupPostgres() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbName,
		dbPassword)

	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func setupSQLite() (*gorm.DB, error) {
	dbLocation := os.Getenv("DATABASE_PATH")
	if dbLocation == "" {
		dbLocation = "/opt/auth-service/gorm.db"
	}

	// Create the sqlite file if it's not available
	if _, err := os.Stat(dbLocation); err != nil {
		if _, err = os.Create(dbLocation); err != nil {
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})
	return db, err
}

func InitializeDatabaseLayer() error {

	dbs := os.Getenv("DB")
	var db *gorm.DB
	var err error

	switch dbs {
	case "sqlite":
		db, err = setupSQLite()
		break
	case "postgres":
		db, err = setupPostgres()
		break
	default:
		return fmt.Errorf("No database found, set the DB env")
	}

	if err != nil {
		return err
	}

	err = models.AutoMigrate(db)
	if err != nil {
		return err
	}
	dbInstance = db
	return nil
}
