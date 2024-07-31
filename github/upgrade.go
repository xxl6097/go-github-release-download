package github

import (
	"errors"
	"fmt"
	"github.com/xxl6097/go-github-publish-release/internal/api"
	"github.com/xxl6097/go-github-publish-release/internal/model"
	"github.com/xxl6097/go-github-publish-release/pkg/dowload"
	"github.com/xxl6097/go-github-publish-release/pkg/version"
	"log/slog"
	"runtime"
	"strings"
)

func checkVersion(curVersion, owner, repo string) (*model.Release, error) {
	release, err := api.GetLatestRelease(owner, repo)
	if err != nil {
		slog.Info("get version error", err)
		return nil, err
	}
	if release == nil {
		return nil, errors.New("release not found")
	}
	if release.TagName == "" {
		return nil, errors.New("TagName is nil")
	}
	v1, v2 := curVersion, release.TagName
	switch version.CompareVersions(v1, v2) {
	case 1:
		fmt.Printf("Version %s is greater than %s.\n", v1, v2)
	case -1:
		fmt.Printf("Version %s is less than %s.\n", v1, v2)
		return release, err
	default:
		fmt.Printf("Version %s and %s are equal.\n", v1, v2)
	}
	return nil, errors.New("no new version available")
}

func Download(dir, name, version, owner, repo string) (string, error) {
	release, err := checkVersion(version, owner, repo)
	if err != nil {
		return "", err
	}
	assets := release.Assets
	if release == nil || assets == nil || len(assets) <= 0 {
		return "", errors.New("release not found")
	}
	ext := ""
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}
	binName := fmt.Sprintf("%s_%s_%s_%s%s", name, release.TagName, runtime.GOOS, runtime.GOARCH, ext)
	var asset *model.Asset
	for _, v := range assets {
		if strings.EqualFold(strings.ToLower(v.Name), strings.ToLower(binName)) {
			asset = &v
		}
	}
	if asset == nil {
		return "", errors.New("asset not found")
	}
	return dowload.DownloadAssets(dir, asset)
}
