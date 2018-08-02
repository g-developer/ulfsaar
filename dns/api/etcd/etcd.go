package etcd


import (
	"context"
	"github.com/coreos/etcd/client"
)

type EtcdClient struct {
	ip string
	port int
	client
	api
}

func NewEtcdClient (ip string, port int) (*EtcdClient, error) {
	url := fmt.Sprintf("http://%v:%v", ip, port)
	cfg := client.Config{
		Endpoints:               []string{url},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	kapi := client.NewKeysAPI(c)
	ins := &EtcdClient{ip, port, c, kapi}
	return ins, nil
}
