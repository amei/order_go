package sqlstore
import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type DbWorker struct {
	//Dsn string
	
}
func init() {
	/*dbw := DbWorker {
		Dsn : "root:123456@tcp(localhost:3306)/order",
	}*/
	db,err := sql.Open("mysql",
	"root:123456@tcp(localhost:3306)/order1",)
	if err != nil {
		fmt.Println("open database error")
		panic(err)
		
	}
	//create table 'orderlist'('id' int(16) NOT NULL AUTO_INCREMENT,'orderid' varchar(256),'ordertype' int(8),'orderuser' varchar(256))

	_,_ = db.Exec(
		"create table orderlist(id int primary key  AUTO_INCREMENT,orderid  varchar(256),ordertype int(8),orderuser varchar(256))")
	defer db.Close()
	
}
func (dbw *DbWorker) Save(orderid,orderuser string,ordertype int) (err error){
	db,err := sql.Open("mysql",
	"root:12345678@tcp(localhost:3306)/order1")
	if err != nil {
		fmt.Println("open database failed")
	}
	_,err = db.Exec(
		"create table orderlist(id int primary key  AUTO_INCREMENT,orderid  varchar(256),ordertype int(8),orderuser varchar(256))")
	if err != nil {
		fmt.Println("create error err",err)
	}
	_, err = db.Exec(
		"INSERT INTO orderlist (orderid, ordertype,orderuser) VALUES (?, ?, ?)",
		orderid,
		ordertype,
		orderuser,
	)
	if err != nil {
		fmt.Println("insert error err",err)
	}
	defer db.Close()
	return err
}