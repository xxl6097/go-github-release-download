package version

import (
	"regexp"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func CompareVersions(version1, version2 string) int {
	// 使用正则表达式分割版本号
	re := regexp.MustCompile(`(\d+)`)
	parts1 := re.FindAllStringSubmatch(version1, -1)
	parts2 := re.FindAllStringSubmatch(version2, -1)

	// 逐级比较版本号
	for i := 0; i < max(len(parts1), len(parts2)); i++ {
		// 将版本号部分转换为整数
		num1, _ := strconv.Atoi(strings.TrimSpace(parts1[i][1]))
		num2, _ := strconv.Atoi(strings.TrimSpace(parts2[i][1]))

		// 比较两个版本号部分
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}

	// 如果所有部分都相等，则版本号相同
	return 0
}
