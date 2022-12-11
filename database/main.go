package database

import (
	"database/sql"
	_ "embed"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

const (
	SEED_DB = "seeds.sqlite"
)

var (
	//go:embed seeds.sqlite
	seedDb []byte

	DB *bun.DB
)

const (
	DATABASE_FILE_PATH = "database.sqlite"
)

func init() {
	err := InitializeDatabase(false)
	if err != nil {
		panic(err)
	}
}

func InitializeDatabase(isReset bool) error {
	if DB != nil {
		DB.Close()
		DB = nil
	}

	if err := tryCopyFile(DATABASE_FILE_PATH, isReset); err != nil {
		return err
	}

	bunDb, err := openDatabase()
	if err != nil {
		return err
	}

	h := bundebug.NewQueryHook(bundebug.WithVerbose(true))
	bunDb.AddQueryHook(h)

	DB = bunDb
	return nil
}

func openDatabase() (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, DATABASE_FILE_PATH)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	return db, nil
}

func tryCopyFile(requestedFilePath string, forceInit bool) error {
	if forceInit {
		file, err := os.Create(requestedFilePath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.Write(seedDb)

		return err
	}

	if _, err := os.Stat(requestedFilePath); err == nil {
		return nil
	}

	file, err := os.Create(requestedFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(seedDb)

	return err
}
