package sqlstore
import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"order_go/projbook/model"
	"time"
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
func (dbw *DbWorker) QueryAll() (data []model.OrderItem) {
	//datamap := make(map[string]model.OrderItem)
	//var datamap []
	db,err := sql.Open("mysql",
	"root:12345678@tcp(localhost:3306)/order1")
	if err != nil {
		fmt.Println("open database failed")
	}
	defer db.Close()
	rows,err := db.Query("select orderid, ordertype,orderuser,createtime from orderlist")
	if err != nil {
		fmt.Println("query error")
	}else {
		for rows.Next() {
			var orderid, ordertype,orderuser,createtime string
			if err = rows.Scan(&orderid,&ordertype,&orderuser,&createtime); err == nil {
				var item = model.OrderItem{}
				item.OrderId = orderid
				item.OrderUser = orderuser
				if ordertype == "order_jd" {
					item.OrderType = "京东订单"
				}else if ordertype == "order_sunning" {
					item.OrderType = "苏宁订单"
				}else if ordertype == "order_mao" {
					item.OrderType = "天猫订单"
				}
				
				item.CreateTime = createtime
				//datamap[orderid] = item
				data = append(data, item)
				fmt.Println("orderid:%s",orderid)
			}
		}
	}
	
	return data
}
func (dbw *DbWorker) Save(orderid,orderuser ,ordertype string) (err error){
	db,err := sql.Open("mysql",
	"root:12345678@tcp(localhost:3306)/order1")
	if err != nil {
		fmt.Println("open database failed")
	}
	defer db.Close()
	_,err = db.Exec(
		"create table if not exists orderlist(id int primary key  AUTO_INCREMENT,orderid  varchar(256),ordertype varchar(128),orderuser varchar(256) ,createtime varchar(256))")
	if err != nil {
		fmt.Println("create error err",err)
	}
	timestr := time.Now().Format("2006-01-02 15:04:05")
	_, err = db.Exec(
		"INSERT INTO orderlist (orderid, ordertype,orderuser,createtime) VALUES (?, ?, ?,?)",
		orderid,
		ordertype,
		orderuser,
		timestr,
	)
	if err != nil {
		fmt.Println("insert error err",err)
	}
	
	return err
}