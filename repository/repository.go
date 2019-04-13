package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/lurenjia528/dockersearch/tags"
	"fmt"
	"os"
)

type Repository struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	RepoName         string `json:"repo_name"`
	ShortDescription string `json:"short_description"`
	StarCount        int    `json:"star_count"`
	PullCount        int    `json:"pull_count"`
	RepoOwner        string `json:"repo_owner"`
	IsAutomated      bool   `json:"is_automated"`
	IsOfficial       bool   `json:"is_official"`
}

func GetRepository(url string, queryRepository string, queryTag string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("网络错误")
		//panic(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("网络错误")
		//panic(err)
		os.Exit(1)
	}
	repositoryResult := Repository{}
	json.Unmarshal([]byte(body), &repositoryResult)

	repositoryKeyWord := strings.Split(queryRepository, ",")
	length := len(repositoryKeyWord)
	result := repositoryResult.Results
	for i := 0; i < len(result); i++ {
		repoName := result[i].RepoName
		for i, keyWord := range repositoryKeyWord {
			if strings.Contains(repoName, keyWord) {
				if i == length-1 {
					if queryTag == "" {
						fmt.Println("详情：" + "https://hub.docker.com/r/" + repoName)
					} else {
						tagOriginUrl := "https://hub.docker.com/v2/repositories/" + repoName + "/tags/?page=1&page_size=250"
						tags.GetTag(tagOriginUrl, repoName, queryTag)
					}
				}
			} else {
				break
			}
		}
	}

	repositoryNextUrl := repositoryResult.Next
	repositoryNextUrl = strings.Replace(repositoryNextUrl, "search-api.s.us-east-1.aws.dckr.io", "hub.docker.com", 1)
	//repositoryNextUrl为空，退出
	if repositoryNextUrl != "" {
		GetRepository(repositoryNextUrl, queryRepository, queryTag)
	}
}
