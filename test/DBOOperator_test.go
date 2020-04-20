package test

import (
	dbo "DBO"
	"fmt"
	"testing"
	"time"
)

func TestNewDBOperator(*testing.T) {
	dbo, err := dbo.NewDBOperator("user:password@tcp(192.168.1.153:3306)/tsdtdb?charset=utf8", "mysql", 30*time.Second, 100, 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dbo)

}
