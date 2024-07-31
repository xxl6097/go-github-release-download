package api

//https://api.github.com/repos/fatedier/frp/releases/latest
import (
	"encoding/json"
	"fmt"
	"github.com/xxl6097/go-github-publish-release/internal/model"
	"log"
	"net/http"
)

// GetReleases 获取项目的发布版本信息
func GetReleases(owner, repo string) ([]model.Releases, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo)

	resp, err := http.Get(url)
	log.Println("resp:", resp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var releases []model.Releases
	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

// GetLatestRelease 获取项目的发布最新版本信息
func GetLatestRelease(owner, repo string) (*model.Release, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	resp, err := http.Get(url)
	//log.Println("resp:", resp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var releases model.Release
	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		return nil, err
	}

	return &releases, nil
}
