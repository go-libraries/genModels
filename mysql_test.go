package genModels

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"github.com/go-libraries/genModels"
)

func TestMysqlGenFile(t *testing.T) {

	mysqlHost := "127.0.0.1"
	mysqlPort := "3306"
	mysqlUser := "root"
	mysqlPassword := "sa"
	mysqlDbname := "mgtj_pay"

	dsn := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDbname + "?charset=utf8mb4"
	fmt.Println(dsn)
	Mysql := genModels.GetMysqlToGo()
	Mysql.Driver.SetDsn(dsn)
	//Mysql.SetStyle("bee")
	Mysql.SetStyle("gorm")
	Mysql.SetModelPath("d:\\Go/src/github.com/go-libraries/genModels/models")
	Mysql.SetIgnoreTables("cate")
	Mysql.SetPackageName("models")
	Mysql.Run()
}



