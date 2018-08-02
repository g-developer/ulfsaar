package main

import (
	"errors"
	"fmt"
	cfg "github.com/g-developer/cfg"
	"io"
	"io/ioutil"
	"net/http"
)

var conf cfg.Config

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	s, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(s))
	io.WriteString(w, string(s) + "\n")
}

func initHttpServer() {
	http.HandleFunc("/gitlab", HelloServer)
	port, err := conf.Get("root.global.port").ToString()
	if nil != err {
		panic(errors.New("Get Listen Port From Cfg Error!"))
	}
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	conf = cfg.Load("../conf/dns.cfg")
	initHttpServer()
}
