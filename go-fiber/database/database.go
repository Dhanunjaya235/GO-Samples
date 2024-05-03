package database

import (
	"database/sql"
	"fmt"
	"gofiber/apis/config"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {

	p := config.GetEnvValueFromKey("DB_PORT")
	port, err := strconv.ParseInt(p, 10, 32)

	if err != nil {
		fmt.Printf("Error Occured While Converting DB_PORT To Integer : %v", err.Error())
		return err
	}

	DB, err = sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.GetEnvValueFromKey("DB_USER"), config.GetEnvValueFromKey("DB_PASSWORD"), config.GetEnvValueFromKey("DB_HOST"), port, config.GetEnvValueFromKey("DB_DATABASE")))

	if err != nil {
		// fmt.Printf("Error While Connecting To MY SQL SERVER : %v", err.Error())
		return err
	}

	if err = DB.Ping(); err != nil {
		// fmt.Printf("Error While Ping : %v", err.Error())
		return err
	}

	DB.Query(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		age INT NOT NULL,
		name VARCHAR(255) NOT NULL
	)
	`)

	fmt.Println("Connected To MY SQL SERVER Successfully")
	return nil

}
