import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
	var err error
	engine,err=xorm.NewEngine("mysql","root:houxi5201314@/todolist_accounts?charset=utf8")
}