package database

import (
	"fmt"
	"log"
	"os"

	"github.com/10Narratives/todo-list/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBInstance represents a database instance.
//
// It contains a pointer to a GORM DB object which is used to
// interact with the database.
type DatabaseInstance struct {
	Db *gorm.DB
}

// DB is a global variable that holds an instance of
// DatabaseInstance, allowing for centralized access
// to the database throughout the application.
var DB DatabaseInstance

// ConnectDB establishes a connection to the PostgreSQL database using
// GORM and runs database migrations for the Task model.
//
// It constructs the database connection string using environment variables:
// DB_USER, DB_PASSWORD, and DB_NAME. If the connection attempt fails,
// it logs an error message and exits the application.
//
// Upon a successful connection, it sets the GORM logger to log at the
// Info level and runs automatic migrations for the Task model.
// After migrating, it assigns the connected database instance to
// a global DatabaseInstance variable called DB.
func ConnectDB() {
	dbUrl := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database.\n", err)
		os.Exit(1)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migration")
	db.AutoMigrate(&models.Task{})

	DB = DatabaseInstance{
		Db: db,
	}
}
