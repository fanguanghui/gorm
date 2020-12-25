package lib

import (
	"database/sql"
	"errors"
	"strconv"
)

func NewDbConnect(dbHost string, dbPort int, dbUser string, dbPassword string, dbDatabase string) (*sql.DB, error) {
	newPassword := ""
	if dbPassword != "" {
		newPassword = ":" + dbPassword
	}
	db, err := sql.Open("mysql", dbUser+newPassword+"@tcp("+dbHost+":"+strconv.Itoa(dbPort)+")/"+dbDatabase+"?&parseTime=True")
	if err != nil {
		return nil, err
	}
	return db, err
}

func GetTablesFromDb(db *sql.DB, dbDatabase string, dbTable string) ([]string, error) {
	var tablesData []string
	if dbTable != "" {
		tablesData = append(tablesData, dbTable)
		return tablesData, nil
	}

	tablesQuery := "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = ?"
	rows, err := db.Query(tablesQuery, dbDatabase)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("No results returned for table\n")
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tablesData = append(tablesData, tableName)
	}
	return tablesData, err
}

func GetColumnsFromDb(db *sql.DB, dbDatabase string, dbTable string) (*map[string]map[string]string, error) {
	columnsData := make(map[string]map[string]string)

	columnsQuery := "SELECT COLUMN_NAME, COLUMN_KEY, DATA_TYPE, IS_NULLABLE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = ? AND table_name = ?"
	rows, err := db.Query(columnsQuery, dbDatabase, dbTable)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("No results returned for table\n")
	}
	defer rows.Close()

	for rows.Next() {
		var column string
		var columnKey string
		var dataType string
		var nullable string
		err = rows.Scan(&column, &columnKey, &dataType, &nullable)
		if err != nil {
			return nil, err
		}
		columnsData[column] = map[string]string{"value": dataType, "nullable": nullable, "primary": columnKey}
	}
	return &columnsData, err
}
