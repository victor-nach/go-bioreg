package config

import (
	"log"

	dbModel "go-bioreg/server/models/db"

	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

func Migrate() {
	connectionString := "postgres://pdknfpiq:J4YO3Ot-szn4REU6DifqmsU610a5ga4l@raja.db.elephantsql.com:5432/pdknfpiq"

		// open a connection to postgres using the gorm library
		db, err := gorm.Open("postgres", connectionString)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Succesfully connected to the db")

		// close connection after the call to the database connection
		defer db.Close()

		// Auto migrate the db
		db.AutoMigrate(&dbModel.User{})
}

// GetConn gets and returns a connection to the database
func GetConn() *gorm.DB {
	connectionString := "postgres://pdknfpiq:J4YO3Ot-szn4REU6DifqmsU610a5ga4l@raja.db.elephantsql.com:5432/pdknfpiq"

	// open a connection to postgres using the gorm library
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Succesfully connected to the db")

	// close connection after the call to the database connection
	// defer db.Close()

	return db
}