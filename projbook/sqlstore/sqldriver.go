package sqlstore
import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type DbWorker struct {
	Dsn string
	
}
func init() {
	dbw := DbWorker {
		Dsn : "root:123456@tcp(localhost:3306)/order",
	}
	db,err := sql.Open("mysql",
	dbw.Dsn)
	if err != nil {
		fmt.Println("open database error")
		panic(err)
		
	}
	_,_ = db.Exec(
		"create table `orderlist`(`id` bigint(20) NOT NULL AUTO_INCREMENT,`orderid` varchar(256),`ordertype` int(8),`orderuser` varchar(256)")
	defer db.Close()
	
}
func (dbw *DbWorker) Save(orderid,orderuser string,ordertype int) (err error){
	db,err := sql.Open("mysql",
	"root:123456@tcp(localhost:3306)/order")
	_, err = db.Exec(
		"INSERT INTO orderlist (orderid, ordertype,orderuser) VALUES (?, ?, ?)",
		orderid,
		ordertype,
		orderuser,
	)
	if err != nil {
		fmt.Println("insert error")
	}
	return err
}