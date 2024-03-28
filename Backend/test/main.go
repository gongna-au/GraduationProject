package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 访问Python服务
		resp, err := http.Get("http://python-service:5000/")
		if err != nil {
			fmt.Fprintf(w, "Error accessing Python service: %s", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(w, "Error reading Python response: %s", err)
			return
		}

		// 连接MySQL数据库
		db, err := sql.Open("mysql", "username:password@protocol(address)/dbname")
		if err != nil {
			fmt.Fprintf(w, "Error connecting to database: %s", err)
			return
		}
		defer db.Close()

		fmt.Fprintf(w, "Python service said: %s\n", string(body))
		// 这里添加数据库操作逻辑
	})

	fmt.Println("Go client is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
