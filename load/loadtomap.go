package load

import (
	"io"
	"os"
	"bufio"
	"strings"
	"regexp"
)

//读取文件
func readfile(url string) ([]string, error) {
	f, err := os.Open(url)

	defer f.Close()

	if err != nil {
		return nil, err
	}

	buf := bufio.NewReader(f)
	propertiesString := make([]string, 0)

	for {
		line, err := buf.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		line = strings.TrimSpace(line)

		//跳过空字符串
		if line == "" {
			continue
		}

		//跳过注释
		match,err:=regexp.MatchString("^#",line)

		if err != nil {
			return nil, err
		}

		if match {
			continue
		}

		propertiesString = append(propertiesString, line)
	}

	return propertiesString, nil
}

//加载properties为map
func LoadProperties(url string) (map[string]string, error) {
	propertiesString, err := readfile(url)

	if err != nil {
		return nil, err
	}

	return convertToMap(propertiesString)
}

//转换为map
func convertToMap(propertiesString []string) (map[string]string, error) {
	propertiesMap := make(map[string]string)

	for _, value := range propertiesString {
		keyvalue := strings.Split(value, "=")

		if len(keyvalue) == 2 {
			propertiesMap[keyvalue[0]] = keyvalue[1]
		}

	}

	return propertiesMap, nil
}
