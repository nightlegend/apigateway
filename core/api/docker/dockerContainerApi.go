package docker

import (
	"container/list"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"strings"
)

func GetAllContainer() *list.List {
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("http://localhost:2375", "v1.21", nil, defaultHeaders)
	if err != nil {
		panic(err)
	}
	//获取本机所有的容器ContainerList(context.Background(), options)
	options := types.ContainerListOptions{All: true}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}
	//打印容器id
	containerList := list.New()
	for _, c := range containers {
		container := make(map[string]string)
		container["id"] = c.ID
		container["image"] = c.Image
		container["status"] = c.Status
		container["command"] = c.Command
		container["name"] = strings.Join(c.Names[:], ",")
		containerList.PushBack(container)
	}
	return containerList
}
