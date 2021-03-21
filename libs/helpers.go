package libs

import (
	"bytes"
	"errors"
	"log"
	"os"
)

func GetPwd() string {
	mydir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can't get dir. error: %v\n", err)
	}
	return mydir
}

func GetDCFile() (string, error) {
	var rootPath bytes.Buffer
	DCFile := "/docker-compose.yml"
	rootPath.WriteString(GetPwd())
	rootPath.WriteString(DCFile)
	if FileExists(rootPath.String()) {
		return rootPath.String(), nil
	}

	var provisionPath bytes.Buffer
	provisionPath.WriteString(GetPwd())
	provisionPath.WriteString("/provisioning")
	provisionPath.WriteString(DCFile)
	if FileExists(provisionPath.String()) {
		return provisionPath.String(), nil
	}
	return "", errors.New("docker-compose file not found")
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func InsertToSlice(slice []string, target string, index int) []string {
	return append(slice[:index], append([]string{target}, slice[index:]...)...)
}
