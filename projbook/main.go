package main

import (
    "fmt"
    "net/http"
    //"strings"
    "log"
    "html/template"
   // "io/ioutil"
    //sqlstore "order_go/projbook/sqlstore"
)
type orderItem struct {
    order_id string
    order_type string
    order_person string
    order_time string
}
func OrderList(w http.ResponseWriter,r *http.Request) {
    fmt.Println("OrderList method ",r.Method);
    if r.Method == "GET" {
        t, _ := template.ParseFiles("./views/orderinfo.gtpl")
       // var orderlist = make(map[string]interface{},1)
     /* orderlist :=[]orderItem {
      orderItem{"11111","order_jd","chenmei","20190107"},
       orderItem{"222222","order_jd","chenmei","20190107"},
       orderItem{"333333","order_jd","chenmei","20190107"},
    }*/
       
        var orderlist [2]orderItem
        orderlist[0].order_id="1111111"
        orderlist[0].order_type="order_jd"
        orderlist[0].order_person = "chenmei"
        orderlist[0].order_time = "20190107"

        orderlist[1].order_id="22222222"
        orderlist[1].order_type="order_jd"
        orderlist[1].order_person = "chenmei"
        orderlist[1].order_time = "20190107"
        
        
        t.Execute(w,orderlist)
    }else {
        r.ParseForm()
        orderId := r.PostForm["orderId"]
        ordertype := r.PostForm["ordertype"]
        order_person := r.PostForm["order_person"]

        fmt.Println("submit body:",orderId,ordertype,order_person)
       // dbw := &sqlstore.DbWorker{}
        //dbw.Save(orderId,order_person,ordertype)
    }
}
func Login(w http.ResponseWriter, r *http.Request) {
    
    if r.Method == "GET" {
        t, _ := template.ParseFiles("./views/login.gtpl")
        t.Execute(w, nil)
    } else {
        //请求的是登陆数据，那么执行登陆的逻辑判断
        r.ParseForm()
        username := r.PostForm["username"]
        password := r.PostForm["password"]
        if username[0] == "order_admin" && password[0] == "admin@123" {
           // w.Write("")
           fmt.Println("login success")
           http.Redirect(w, r, "/list", http.StatusFound)
           
        }
        
        
        //fmt.Println("username:", r.PostForm["username"])
        //fmt.Println("password:", r.PostForm["password"])
    }
}

func main() {
    http.HandleFunc("/login", Login) //设置访问的路由
    http.HandleFunc("/list",OrderList)
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}