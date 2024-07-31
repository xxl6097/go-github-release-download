package dowload

import (
	"github.com/xxl6097/go-github-publish-release/internal/model"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadAssets 下载发布版本的文件
func DownloadAssets(dir string, asset *model.Asset) (string, error) {
	log.Println("Downloading...", asset.BrowserDownloadUrl)
	resp, err := http.Get(asset.BrowserDownloadUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	newFilePath := filepath.Join(dir, asset.Name)
	// 创建文件
	file, err := os.Create(newFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 将响应体写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}
	return newFilePath, nil
}
