package data

import (
	"cyberuranus-interview/pkg/common"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	selectSQL = `
	SELECT id, lat, lng, date_added
	FROM data;`

	insertSQL = `
	INSERT INTO data(id, lat, lng, date_added)
	VALUES(?, ?, ?, ?);`
)

// Data represents the specific structure from CyberUranus
type Data struct {
	ID       string `json:"id" form:"id"`
	Location struct {
		Lat  float32 `json:"lat" form:"lat"`
		Long float32 `json:"long" form:"lng"`
	} `json:"location" form:"location"`
	DateAdded time.Time `json:"data_added,omitempty" form:"date_added"`
}

// GetData returns all of data from db, if error exists, return the error only.
func GetData(db *sql.DB) ([]*Data, error) {
	rows, err := db.Query(selectSQL)
	if err != nil {
		return nil, fmt.Errorf("Get Data Error: %w", err)
	}

	allData := make([]*Data, 0)
	for rows.Next() {
		var dataAddedTime int64

		var d Data
		if err := rows.Scan(&d.ID, &d.Location.Lat, &d.Location.Long, &dataAddedTime); err != nil {
			return nil, fmt.Errorf("Get Data Error: %w", err)
		}
		d.DateAdded = time.Unix(dataAddedTime, 0)
		allData = append(allData, &d)
	}

	return allData, nil
}

// InsertData insert a data to db, if error exists, return the error only.
func InsertData(db *sql.DB, d Data) error {
	dataAddedSQL := d.DateAdded.Unix()

	if _, err := db.Exec(insertSQL, d.ID, d.Location.Lat, d.Location.Long, dataAddedSQL); err != nil {
		// TODO: Should be handle difference mysql error situations to the client.
		// Handle Duplicate Entry only for current version
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == common.MysqlErrDupEntry {
				return errors.New("Duplicated ID")
			}
		}

		return fmt.Errorf("Insert Data Error: %w", err)
	}

	return nil
}
