package github

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GithubRepository struct {
	Forks    int32  `json:"forks_count"`
	Name     string `json:"name"`
	Stars    int32  `json:"stargazers_count"`
	Watchers int32  `json:"watchers_count"`
}

func (client *GithubClient) GetRepository(name string) (GithubRepository, error) {
	var result GithubRepository
	url := "https://api.github.com/repos/" + name
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "application/vnd.github.v3+json")

	if err != nil {
		return result, err
	}

	resp, err := client.http_client.Do(request)

	if err != nil {
		return result, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		err = json.Unmarshal(body, &result)
		if err != nil {
			return result, nil
		}
	} else {
		return result, errors.New("repository your search doesn't exsits")
	}

	return result, nil
}

func (client *GithubClient) ProfileExists(name string) bool {
	repo, err := client.GetRepository(name + "/" + name)
	if err != nil {
		return false
	}
	if repo.Name != "" {
		return true
	}
	return false
}
