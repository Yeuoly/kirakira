package github

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"
)

type GithubRepository struct {
	Forks    int32  `json:"forks_count"`
	Name     string `json:"name"`
	Stars    int32  `json:"stargazers_count"`
	Watchers int32  `json:"watchers_count"`
}

func getUserAgent() []string {
	return []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36 Edg/91.0.864.54",
		"User-Agent:Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"User-Agent:Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"User-Agent: Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent:Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}
}

func (client *GithubClient) GetRepository(name string) (GithubRepository, error) {
	user_agents := getUserAgent()

	var result GithubRepository
	url := "https://api.github.com/repos/" + name
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "application/vnd.github.v3+json")

	random, _ := rand.Int(rand.Reader, big.NewInt(4))
	ua := user_agents[random.Int64()]
	request.Header.Set("UserAgent", ua)
	request.Header.Set("Authorization", "token "+client.access_token)

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
			return result, err
		}
	} else {
		return result, errors.New("repository your search doesn't exsits or network error")
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
