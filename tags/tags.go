package tags

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Tags struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name        string   `json:"name"`
	FullSize    int      `json:"full_size"`
	Images      []Images `json:"images"`
	Id          int      `json:"id"`
	Repository  int      `json:"repository"`
	Creator     int      `json:"creator"`
	LastUpdater int      `json:"last_updater"`
	LastUpdated string   `json:"last_updated"`
	ImageId     string   `json:"image_id"`
	V2          bool     `json:"v2"`
}

type Images struct {
	Size         int    `json:"size"`
	Architecture string `json:"architecture"`
	Variant      string `json:"variant"`
	Features     string `json:"features"`
	Os           string `json:"os"`
	OsVersion    string `json:"os_version"`
	OsFeatures   string `json:"os_features"`
}

func GetTag(url string, repoName string, archTag string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	tagResult := Tags{}
	json.Unmarshal([]byte(body), &tagResult)

	tagNextUrl := tagResult.Next

	result := tagResult.Results
	for i := 0; i < len(result); i++ {
		tagName := result[i].Name
		if strings.Contains(tagName, archTag) {
			repositoryAndTag := repoName + ":" + tagName
//			fmt.Println(repositoryAndTag)
			fmt.Println("详情：" + "https://hub.docker.com/r/" + repoName + "/tags/")
			fmt.Println("下载：" + "docker pull " + repositoryAndTag)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			continue
		} else {
			images := result[i].Images

			for j := 0; j < len(images); j++ {
				arch := images[j].Architecture
				if strings.Contains(arch, archTag) {
					repositoryAndTag := repoName + ":" + tagName
//					fmt.Println(repositoryAndTag)
					fmt.Println("详情：" + "https://hub.docker.com/r/" + repoName + "/tags/")
					fmt.Println("下载：" + "docker pull " + repositoryAndTag)
					fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
				}
			}
		}
	}

	if tagNextUrl != "" {
		GetTag(tagNextUrl, repoName, archTag)
	}
}
