package str

import (
	"os"
	"regexp"
	"strings"
)

// ReplaceEnvVariables 查找字符串中的环境变量并替换为其值
func ReplaceEnvVariables(input string) string {
	// 定义正则表达式以匹配 ${ENV_VAR} 格式的环境变量
	envVariableRegex := regexp.MustCompile(`\$\{([^\}]+)\}`)

	// 替换函数，根据匹配到的环境变量获取其值
	replacer := func(match string) string {
		envVariableName := strings.Trim(match, "${}")
		envVariableValue := os.Getenv(envVariableName)
		return envVariableValue
	}

	// 使用替换函数替换所有匹配到的环境变量
	output := envVariableRegex.ReplaceAllStringFunc(input, replacer)

	return output
}

// FindEnvVariables 查找字符串中的环境变量
func FindEnvVariables(input string) map[string]string {
	envVariableRegex := regexp.MustCompile(`\$\{([^\}]+)\}`) // 正则表达式匹配 ${ENV_VAR} 格式的字符串
	matches := envVariableRegex.FindAllStringSubmatch(input, -1)

	envVariables := make(map[string]string)

	for _, match := range matches {
		envVariableName := match[1]
		envVariableValue := os.Getenv(envVariableName)

		if envVariableValue != "" {
			envVariables[envVariableName] = envVariableValue
		}
	}

	return envVariables
}
