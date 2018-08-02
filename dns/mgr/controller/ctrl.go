package ctl

import (
	util "github.com/g-developer/util/go"
	"sync"
)

type ctl struct {
	port int
	applyed map[string]bool
	mutex *sync.Mutex
	
}

type ApiDesc struct {
	uri string
	handler util.Handler
}

func RegisterApi (name string, api ApiDesc) {
	if nil == ins {
		painc("Must NewServer First!")
	}
	if exist, err := ins.applyed[name]; ok {
		return
	} else {
		ins.mutex.Lock()
		err := hs.AddHandler(api.uri, api.handler)
		if nil != err {
			fmt.Println("add ", api.uri, " error!")
		} else {
			ins.applyed[name] = true
		}
		ins.mutex.Unlock()
	}
}

var once sync.Once
var ins *ctl
var hs := util.GetHttpMgrInstance()

func NewServer (int port) *ctl {
	once.Do(func() {
		hs.NewHttpServer(port)
		ins = &ctr{int, map[string]bool{}, &sync.Mutex{}}
	})
	return ins
}

