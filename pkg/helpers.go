package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func GetPwd() string {
	mydir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return mydir
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

func ReplaceInternalVariables(from string, to string, list []string) []string {
	for i, v := range list {
		r := regexp.MustCompile(from)
		if found := r.FindAllString(v, -1); found != nil {
			for _, vv := range found {
				list[i] = strings.Replace(list[i], vv, to, -1)
			}
		}
	}
	return list
}

func ReplaceEnvVariables(args []string) []string {
	for i, v := range args {
		r := regexp.MustCompile("{(.*?)}")
		if found := r.FindAllString(v, -1); found != nil {
			for _, vv := range found {
				args[i] = replaceEnv(args[i], vv)
			}
		}
	}

	return args
}

func replaceEnv(target, pattern string) string {
	trimmed := strings.TrimRight(strings.TrimLeft(pattern, "{"), "}")
	if ev := os.Getenv(trimmed); ev != "" {
		return strings.Replace(target, pattern, ev, 1)
	}
	return target
}

func MergeSliceOfString(data []string) []string {
	return DeleteEmpty(strings.Split(strings.Join(data, " "), " "))
}

func IndexExists(slice []string, index int) bool {
	if (len(slice) - 1) >= index {
		return true
	}
	return false
}

func DockerExec(signature []string) []byte {
	cmd := exec.Command("docker", signature...)
	data, err := cmd.Output()
	if err != nil {
		log.Fatalf("Cant exec docker command, err: %s", err)
	}

	return data
}

func RetrieveGatewayHost(data []byte) string {
	var networkInspected []NetworkInspected
	if err := json.Unmarshal(data, &networkInspected); err != nil {
		log.Fatal(err)
	}

	if len(networkInspected) == 0 {
		log.Println("cant network inspect")
		return ""
	}

	return networkInspected[0].GetGateway()
}

func DownloadFile(uri, fileName, dst string) string {
	tmpFile := fmt.Sprintf("%s/%s", dst, fileName)
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalf("error while get latest file, err: %s\n", err)
		return ""
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("error while close response body, err: %s\n", err)
		}
	}()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error while read source, err: %s\n", err)
		return ""
	}

	if err := ioutil.WriteFile(tmpFile, data, 0755); err != nil {
		log.Fatalf("error while write file, err: %s\n", err)
		return ""
	}

	return tmpFile
}

func FlattenSlice(target [][]string) []string {
	var flatten []string

	for _, v := range target {
		flatten = append(flatten, v...)
	}

	return flatten
}

func Contains(data []string, needle string) bool {
	for _, v := range data {
		if needle == v {
			return true
		}
	}
	return false
}
