package main

import (
	"encoding/json"

	"github.com/swanwish/go-common/logs"
	"github.com/swanwish/go-common/utils"
)

func main() {
	url := `https://s.natalian.org/2016-12-07/debugme2.json`
	_, content, err := utils.GetUrlContent(url)
	if err != nil {
		logs.Errorf("Failed to get content from url %s, the error is %v", url, err)
		return
	}

	movies := []struct {
		Title       string `json:"title"`
		PublishedAt string `json:"published_at"`
	}{}
	err = json.Unmarshal(content, &movies)
	if err != nil {
		logs.Errorf("Failed to unmarshal content %s, the error is %v", string(content), err)
		return
	}
	logs.Debugf("The movies are %v", movies)
}
