package core

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strconv"

	"github.com/Yeuoly/kirakira/github"
)

func ReplaceMD(client *github.GithubClient, filename string) (string, error) {
	regx := regexp.MustCompile("{.+}\\(github\\.com\\/([a-zA-Z0-9_]+)\\/([a-zA-Z0-9_]+)\\)")

	regx_reponame := regexp.MustCompile("\\/(([a-zA-Z0-9_]+)\\/([a-zA-Z0-9_]+))")
	regx_stars := regexp.MustCompile("\\(__stars__\\)")
	regx_forks := regexp.MustCompile("\\(__forks__\\)")
	regx_group := regexp.MustCompile("\\{(.+)\\}")

	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", errors.New("不存在的文件")
	}

	//var temp map[string]github.GithubRepository
	//temp = make(map[string]github.GithubRepository)

	file = regx.ReplaceAllFunc(file, func(b []byte) []byte {
		//匹配仓库名
		reponames := regx_reponame.FindSubmatch(b)
		var reponame []byte
		if len(reponames) > 0 {
			reponame = reponames[1]
		} else {
			return b
		}
		repo, err := client.GetRepository(string(reponame))

		if err != nil {
			return b
		}

		b = regx_stars.ReplaceAllFunc(b, func(c []byte) []byte {
			return []byte(strconv.Itoa(int(repo.Stars)))
		})

		b = regx_forks.ReplaceAllFunc(b, func(c []byte) []byte {
			return []byte(strconv.Itoa(int(repo.Forks)))
		})

		return regx_group.FindSubmatch(b)[1]
	})

	return string(file), nil
}
