package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var config map[string]string = make(map[string]string, 0)

func LoadConfig(key string) (string, error) {
	// str := config[key]
	fmt.Printf("config:%v", config)
	value := config[key]

	return value, nil
}

func init() {
	fp, err := os.Open("/workspace/go/.env")
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	defer fp.Close()

	var ioReader = bufio.NewReaderSize(fp, 10)
	for {
		lineStr, err := ReadLine(ioReader, "")
		if err != nil {
			fmt.Printf("err:%v\n", err)
			return
		}
		// fmt.Printf("line:%v\n", lineStr)
		strs := strings.Split(lineStr, "=")
		config[strs[0]] = strs[1]
	}

}

func ReadLine(ioReader *bufio.Reader, preLine string) (string, error) {
	line, isPrefix, err := ioReader.ReadLine()
	if err != nil {
		fmt.Printf("err:%v", err)
		return "", err
	}
	lineStr := string(line)
	lineStr = fmt.Sprintf("%s%s", preLine, lineStr)
	if isPrefix {
		// fmt.Printf("full\n")
		return ReadLine(ioReader, lineStr)
	}

	return lineStr, nil
}
