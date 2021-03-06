package database

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"library/structs"
	"log"
	"os"
	"strings"
	"time"
)

var readDB *sql.DB

func LoadDatabaseConfig(fileName string) {
	configFile, err := os.Open(fileName)
	defer configFile.Close()

	settings := structs.DatabaseConfig{}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&settings)

	if err != nil {
		log.Fatal(`Fatal Error: Unable to load database configuration!
						Looking for file : `, fileName)
	}

	structs.Database = settings
}

func GetConnectionString() string {
	var connectionString strings.Builder
	connectionString.WriteString(structs.Database.User)
	connectionString.WriteString(":")
	connectionString.WriteString(structs.Database.Password)
	connectionString.WriteString("@tcp(")
	connectionString.WriteString(structs.Database.Url)
	connectionString.WriteString(":")
	connectionString.WriteString(structs.Database.Port)
	connectionString.WriteString(")/")
	connectionString.WriteString(structs.Database.DbName)
	connectionString.WriteString("?parseTime=true")
	return connectionString.String()
}

func OpenGetSqlDB() *sql.DB {
	openString := GetConnectionString()

	db, err := sql.Open("mysql", openString)
	if err != nil {
		log.Fatal("NO DATABASE!", err)
	}
	db.SetConnMaxLifetime(5 * time.Minute)

	readDB = db
	return db
}

func getSqlDB() *sql.DB {
	if readDB != nil {
		return readDB
	}
	return OpenGetSqlDB()
}

func GetSqlReadDB() *sql.DB {
	return getSqlDB()
}

// One day we might write to a different db
func GetSqlWriteDB() *sql.DB {
	return getSqlDB()
}
