package main

import (
    "fmt"
    "net/http"
    //"strings"
    "log"
    "html/template"
   // "io/ioutil"
    sqlstore "order_go/projbook/sqlstore"
    //"order_go/projbook/model"
)
/*type OrderItem struct {
    Order_id string
    Order_type string
    Order_person string
    Order_time string
}*/
func OrderList(w http.ResponseWriter,r *http.Request) {
    fmt.Println("OrderList method ",r.Method);
    if r.Method == "GET" {
        t, _ := template.ParseFiles("./views/orderinfo.gtpl")
        dbw := &sqlstore.DbWorker{}
     
        data := dbw.QueryAll()

        
        err := t.Execute(w,data)
        fmt.Println("-----------------------",err)
    }else {
        r.ParseForm()
        fmt.Println("formid = ",r.Form["id"])
        orderId := r.PostForm["orderId"]
        ordertype := r.PostForm["ordertype"]
        order_person := r.PostForm["order_person"]

        fmt.Println("submit body:",orderId,ordertype,order_person)
        dbw := &sqlstore.DbWorker{}
        dbw.Save(orderId[0],order_person[0],ordertype[0])
        t, _ := template.ParseFiles("./views/orderinfo.gtpl")
     
        data := dbw.QueryAll()
       
        err := t.Execute(w,data)
        fmt.Println("-----------------------",err)
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