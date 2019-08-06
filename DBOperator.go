package dbo

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBOperator struct {
	Db *sql.DB
}

/**
* 执行select查询
*
 */
func (dbo *DBOperator) Query(sqlStr string, params ...interface{}) ([]map[string]sql.RawBytes, error) {
	//log.Printf("query sql:%s,params:%+v", sqlStr, params)
	result := make([]map[string]sql.RawBytes, 5)
	result = result[0:0]
	stmt, err := dbo.Db.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(params...)
	if err != nil {
		log.Fatal(err)
		return result, err
	}

	fields, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
		return result, err
	}

	values := make([]sql.RawBytes, len(fields))
	scanArgs := make([]interface{}, len(fields))

	for i := range scanArgs {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
			continue
		}

		rowData := make(map[string]sql.RawBytes)
		for i := 0; i < len(fields); i++ {
			rowData[fields[i]] = values[i]
		}
		result = append(result, rowData)
	}
	return result, nil
}

/**
* 执行update insert delete等操作
*
 */
func (dbo *DBOperator) Execute(sqlStr string, params ...interface{}) (int64, error) {
	//log.Printf("execute sql:%s,params:%+v", sqlStr, params)
	stmt, err := dbo.Db.Prepare(sqlStr)
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, err
	}
	num, err := result.RowsAffected()
	return num, err
}