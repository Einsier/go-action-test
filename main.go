package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/", HelloServer) //设置server处理函数
	fmt.Println("Server is going to start..")
	err := http.ListenAndServe(":9090", nil) //设置监听的端口，并启动服务
	fmt.Println("Server is going to quit..")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
