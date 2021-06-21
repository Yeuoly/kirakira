package github

import (
	"net/http"
	"net/url"
)

type GithubClient struct {
	http_client  *http.Client
	access_token string
}

func (client *GithubClient) Init(proxy_url string, access_token string) error {
	client.access_token = access_token

	if proxy_url != "" {
		urli := url.URL{}
		proxyurl, err := urli.Parse(proxy_url)

		if err != nil {
			return err
		}

		client.http_client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyurl),
			},
		}
	} else {
		client.http_client = &http.Client{}
	}
	return nil
}
