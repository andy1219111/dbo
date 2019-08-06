# dbo
go数据库操作的封装

##使用

```golang
import (
	"github.com/andy1219111/dbo"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//得到数据库连接对象
dbDataCenter, err := sql.Open(configObj.Database["DBDriver"], configObj.Database["data_center"])
defer dbDataCenter.Close()
dbOperator := dbo.DBOperator{Db: dbDataCenter}
```

#### 查询

```
userID := 15

//查询
sqlStr := `
		SELECT
		*
		FROM
		user
		WHERE user_id = ?

`
//result为包含查询结果的数组  数据每个元素为map[string]sql.RawBytes,只需要进行相应的数据类型转换即可
result, err := dbOperator.Query(sqlStr, userID)
if err != nil {
	log.Printf("sql query error:%+v", err)
}
userName := string(result[0]["user_name"])
age, _ := strconv.Atoi(string(result[0]["age"]))

```

#### 执行update、insert等操作

```
sqlStr := `
		INSERT INTO user (
			user_name,
			age
		  )
		VALUES
			(?,?)

`
//入库
rowAffected, err := dbOperator.Execute(sqlStr, "aaron", "15")
if err != nil {
	log.Printf("更新扫码数据失败:%+v", err)
}

log.Println("插入用户信息成功")
```
