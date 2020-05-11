package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/go/src/pkg/strconv"
	"github.com/micro/go-micro/cmd"
	"go/go-micro/demo/src/share/config"
	"go/go-micro/demo/src/share/utils/path"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/",handlerRPC)
	log.Println("Listen on:8888")
	e := http.ListenAndServe(":8888", mux)
	if e !=nil{
		fmt.Println(e)
	}

}

func handlerRPC(w http.ResponseWriter, r *http.Request)  {
	log.Println("handlerRPC...")
	// 1.正常请求
	if r.URL.Path == "/"{
		_, e := w.Write([]byte("server..."))
		if e!=nil{
			fmt.Println(e)
		}
		return
	}
	// 2.RPC请求，跨域请求
	if origin := r.Header.Get("Origin");true{
		w.Header().Set("Access-Control-Allow-Origin",origin)
	}
	w.Header().Set("Access-Control-Allow-Methods","POST GET")
	w.Header().Set("Access-Control-Allow-Headers","X-Token,Content-Type,Content-Length")
	w.Header().Set("Access-Control-Allow-Credentials","true")
	handlerJSONRPC(w, r)
	return
}

func handlerJSONRPC(w http.ResponseWriter,r*http.Request)  {
	service, method := path.PathToReceiver(config.Namespace, r.URL.Path)
	log.Println("service:" + service)
	log.Println("method:" + method)
	// 读取请求体
	br, _ := ioutil.ReadAll(r.Body)
	request := json.RawMessage(br)
	var response json.RawMessage
	req := (*cmd.DefaultOptions().Client).NewJsonRequest(service,method,&request)
	ctx := path.RequestToContext(r)
	err := (*cmd.DefaultOptions().Client).Call(ctx, req,&response)
	if err != nil{
		fmt.Println(err)
		return
	}
	b, _ := response.MarshalJSON()
	// 设置响应头
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	_, err = w.Write(b)
	if err !=nil{
		fmt.Println(err)
	}
}
