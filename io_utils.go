package go_commons

import (
	"bufio"
	"io/ioutil"
	"os"
)

type ProcessTextLine func(string) (interface{}, error)

func ReadRawFile(file string) (string, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return "", err
	}

	b, err := ioutil.ReadFile(file)
	return string(b), err
}

func WriteTextFile(file string, content string) error {
	return ioutil.WriteFile(file, []byte(content), 0644)
}

func ReadTextLineByLine(file string, p ProcessTextLine) ([]interface{}, error) {
	result := make([]interface{}, 0)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return result, err
	}

	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return result, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		r, err := p(line)
		if err != nil {
			continue
		}
		if r != nil {
			result = append(result, r)
		}
	}

	return result, nil
}
