// Package etcd provides the etcd backend plugin.
package etcd

import (
	"fmt"
	"strings"
	"github.com/coredns/coredns/plugin/etcd/msg"
	"github.com/coredns/coredns/request"
	"github.com/coredns/coredns/plugin/test"
)

func (e *Etcd) isInExcept (name string) bool {
	//check in etcd
	path := e.Except + "/" + name
	r, err := e.get(path, true)
	fmt.Printf("record--------%v; err---%v\n", r, err)
	if nil == err {
		return true
	} else {
		return false
	}
}

func (e *Etcd) except (state request.Request, exact bool) ([]msg.Service, error) {
	sx := []msg.Service{}
	state.W = &test.ResponseWriter{}
       	resp, err := e.Upstream.Forward.Forward(state)
	if nil == err {
		if nil != resp.Answer {
			lens := len(resp.Answer)
			for i:=0; i<lens; i++ {
				serv := new(msg.Service)
				str := resp.Answer[i].String()
				fmt.Printf("resp.Answer[%v]=[%#v]\n", i, str)
				arr := strings.Split(str, "\t")
				arrLen := len(arr)
				ip := arr[arrLen - 1]
				serv.Host = ip
				fmt.Printf("server--333------%v\n", ip)
				fmt.Printf("server--444------%+v\n", *serv)
				sx = append(sx, *serv)
			}
		}
		fmt.Printf("finally---------count(sx) = %v\n", len(sx))
		return sx, nil
	} else {
		return nil, err
	}
}

