package config

import (
	"os"

	"github.com/cplxx/goppurtunities.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	//check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err){
		logger.Info("database file not found, creating")
		//Create the database file and directory 
		err = os.Mkdir("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
			if err != nil {
				return nil, err
		}
		file.Close()
	}

	//create db and connect
	db, err	:= gorm.Open(sqlite.Open(dbPath), &gorm.Config{} )
	if err != nil {
		logger.ErrorF("sqlite opening error: %v", err)
		return nil, err 
	}

	//Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("sqlite automigration error: %v", err)
		return nil, err
	}

	//Return the db
	return db, nil
}
