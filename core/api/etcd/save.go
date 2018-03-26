package etcd

import (
	"github.com/nightlegend/apigateway/core/utils/etcd"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// SaveAction :save data to etcd.
func SaveAction() {
	kapi := etcd.EtcdConn()
	resp, err := kapi.Set(context.Background(), "/test", "test", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
	// kapi.Watcher(key, opts)

}
