package etcd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/nightlegend/apigateway/core/utils/etcd"
	"golang.org/x/net/context"
)

// SaveAction :save data to etcd.
func SaveAction() {
	kapi := etcd.Conn()
	resp, err := kapi.Set(context.Background(), "/test", "test", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
	// kapi.Watcher(key, opts)

}
