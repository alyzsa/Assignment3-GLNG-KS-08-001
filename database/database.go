package database

import (
    "database/sql"
    "errors"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func verifyDBConnection() error {
    if DB == nil {
        return errors.New("The database connection has not been initialized")
    }
    if err := DB.Ping(); err != nil {
        return fmt.Errorf("Database connection failed: %v", err)
    }
    return nil
}

func InitializeDatabase() error {
    var err error
    DB, err = sql.Open("postgres", "user=postgres password=alysha09 dbname=postgres sslmode=disable")

    if err != nil {
        return err
    }

    if err := verifyDBConnection(); err != nil {
        return err
    }

    log.Println("Database connection established successfully")

    _, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS weather_data(
        id SERIAL PRIMARY KEY, -- Utilizing SERIAL for automated ID generation
        water INT,
        wind INT,
        water_status VARCHAR(255),
        wind_status VARCHAR(255),
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )
    `)

    if err != nil {
        return err
    }
    return nil
}

func CloseDatabase() {
    if DB != nil {
        DB.Close()
    }
}
