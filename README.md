# go-github-release-download
基于github api获取release并下载，可检测版本号


# Usage

```
path, err := github.Download("./", "AuGoService", "0.0.0", "xxl6097", "go-service-framework")
if err == nil {
  log.Println(path)
}
```