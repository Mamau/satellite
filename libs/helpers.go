package libs

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

func GetPwd() string {
	mydir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can't get dir. error: %v\n", err)
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
	return pattern
}

func LoadEnv() {
	if !FileExists(".env") {
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
