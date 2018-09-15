package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"dockersearch/tags"
)

type Repository struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
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
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	repositoryResult := Repository{}
	json.Unmarshal([]byte(body), &repositoryResult)

	result := repositoryResult.Results
	for i := 0; i < len(result); i++ {
		repoName := result[i].RepoName
		if strings.Contains(repoName, queryRepository) {
			tagOriginUrl := "https://hub.docker.com/v2/repositories/" + repoName + "/tags/?page=1&page_size=250"
			tags.GetTag(tagOriginUrl, repoName, queryTag)
		}
	}

	repositoryNextUrl := repositoryResult.Next
	repositoryNextUrl = strings.Replace(repositoryNextUrl, "search-api.s.us-east-1.aws.dckr.io", "hub.docker.com", 1)
	//repositoryNextUrl为空，退出
	if repositoryNextUrl != "" {
		GetRepository(repositoryNextUrl, queryRepository, queryTag)
	}
}