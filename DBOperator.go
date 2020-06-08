package dbo

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//DBOperator 数据库操作对象
type DBOperator struct {
	Db *sql.DB
	//数据源配置
	DSN string
	//数据库驱动类型
	DriverName string
	//连接最大复用时间
	ConnMaxLifetime time.Duration
	//设置最大打开的连接数，默认值为0表示不限制
	MaxOpenConns int
	//设置闲置的连接数
	MaxIdleConns int
}

//NewDBOperator 实例化数据访问对象
func NewDBOperator(dsn, driverName string, ConnMaxLifetime time.Duration, MaxOpenConns, MaxIdleConns int) (*DBOperator, error) {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(ConnMaxLifetime)
	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)
	instance := &DBOperator{
		Db:              db,
		DSN:             dsn,
		DriverName:      driverName,
		ConnMaxLifetime: ConnMaxLifetime,
		MaxOpenConns:    MaxOpenConns,
		MaxIdleConns:    MaxIdleConns,
	}

	return instance, nil
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
	//关闭该连接
	defer rows.Close()
	fields, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
		return result, err
	}

	for rows.Next() {
		values := make([]sql.RawBytes, len(fields))
		scanArgs := make([]interface{}, len(fields))

		for i := 0; i < len(fields); i++ {
			scanArgs[i] = &values[i]
		}
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
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, err
	}

	num, err := result.RowsAffected()
	return num, err
}

//Close 关闭该数据库连接对象
func (dbo *DBOperator) Close() {
	dbo.Db.Close()
}
