package io

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

//usage: lines, err := GetLinesFromFileAsBytes(FILEPATH)
//then can access as string by casting: line := string(lines[idx][:])
func GetLinesFromFileAsBytes(filepath string) ([][]byte, error) {
	file_content_as_bytes, e := ioutil.ReadFile(filepath)
	if e != nil {
		var empty_bytes [][]byte
		empty_bytes = make([][]byte, 0)
		return empty_bytes, fmt.Errorf("Error reading lines from file '%s', %v\n", filepath, e)
	}
	return bytes.Split(file_content_as_bytes, []byte("\n")), nil
}

//usage: var lines []string; err := GetLinesFromFileAsString(FILEPATH, &lines)
func GetLinesFromFileAsString(filepath string, lines *[]string) error {
	file_content_as_bytes, e := ioutil.ReadFile(filepath)
	if e != nil {
		var empty_lines []string
		empty_lines = make([]string, 0)
		*lines = empty_lines
		return fmt.Errorf("Error reading lines from file '%s', %v\n", filepath, e)
	}
	*lines = strings.Split(string(file_content_as_bytes[:]), "\n")
	return nil
}
