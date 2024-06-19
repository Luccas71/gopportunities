package config

import (
	"os"

	"github.com/Luccas71/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	//check if the db file exists
	_, err := os.Stat(dbPath) //verificando se existe o arquivo
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		//create database and file directory
		err = os.MkdirAll("./db", os.ModePerm) //criando o diretorio
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath) //criando o arquivo
		if err != nil {
			return nil, err
		}

		//importante fechar o arquivo!!
		file.Close()
	}

	//create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	//migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	//return the db
	return db, nil
}
