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
type OrderItem struct {
    Order_id string
    Order_type string
    Order_person string
    Order_time string
}
func OrderList(w http.ResponseWriter,r *http.Request) {
    fmt.Println("OrderList method ",r.Method);
    if r.Method == "GET" {
        t, _ := template.ParseFiles("./views/orderinfo.gtpl")
     

        Orderlist := [2]OrderItem{}
        Orderlist[0].Order_id="1111111"
        Orderlist[0].Order_type="order_jd"
        Orderlist[0].Order_person = "chenmei"
        Orderlist[0].Order_time = "20190107"

        Orderlist[1].Order_id="22222222"
        Orderlist[1].Order_type="order_jd"
        Orderlist[1].Order_person = "chenmei"
        Orderlist[1].Order_time = "20190107"
        
        
        err := t.Execute(w,Orderlist)
        fmt.Println("-----------------------",err)
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