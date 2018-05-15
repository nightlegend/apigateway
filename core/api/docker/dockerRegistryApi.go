package docker

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// docker engine info
const (
	ACCOUNT  string = "test"
	PASSWORD string = "test"
	HOST     string = "test"
	CONTENT  string = "v2"
)

// GetRegistryImages is get all images info.
func GetRegistryImages() []byte {
	resp, err := http.Get("http://" + ACCOUNT + ":" + PASSWORD + "@" + HOST + "/" + CONTENT + "/" + "_catalog?n=100&last=a")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body
}

// GetAllTagByImageName is query image by tag name.
func GetAllTagByImageName(imageName string) []byte {
	resp, err := http.Get("http://" + ACCOUNT + ":" + PASSWORD + "@" + HOST + "/" + CONTENT + "/" + imageName + "/tags/list")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body
}

// GetImageTagInfo is get tag info.
func GetImageTagInfo(imageName string, imageTag string) []byte {
	resp, err := http.Get("http://" + ACCOUNT + ":" + PASSWORD + "@" + HOST + "/" + CONTENT + "/" + imageName + "/manifests/" + imageTag)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body
}
