package docker

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ACCOUNT  string = "test"
	PASSWORD string = "test"
	HOST     string = "test"
	CONTENT  string = "v2"
)

func GetRegistryImages() []byte {
	resp, err := http.Get("http://" + ACCOUNT + ":" + PASSWORD + "@" + HOST + "/" + CONTENT + "/" + "_catalog?n=100&last=a")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func GetAllTagByImageName(imageName string) []byte {
	resp, err := http.Get("http://" + ACCOUNT + ":" + PASSWORD + "@" + HOST + "/" + CONTENT + "/" + imageName + "/tags/list")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func GetImageTagInfo(imageName string, imageTag string) []byte {
	resp, err := http.Get("http://" + ACCOUNT + ":" + PASSWORD + "@" + HOST + "/" + CONTENT + "/" + imageName + "/manifests/" + imageTag)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body
}
