package etcd

import (
	"time"

	"github.com/coreos/etcd/client"
	log "github.com/sirupsen/logrus"
)

// EtcdConn :
// connect to etcd
func EtcdConn() client.KeysAPI {
	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	return kapi
}
