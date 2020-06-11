package test

import (
	dbo "DBO"
	"fmt"
	"testing"
	"time"
)

func TestNewDBOperator(*testing.T) {
	dbo, err := dbo.NewDBOperator("user:pwd@tcp(192.168.9.153:3306)/tsdtdb?charset=utf8", "mysql", 30*time.Second, 100, 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dbo)
	results, _ := dbo.Query("select vlan_id,label,name from dt_vlan where deleted!=1")
	for _, one := range results {
		fmt.Println(string(one["vlan_id"]))
	}
}
