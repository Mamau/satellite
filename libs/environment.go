package libs

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func ReplaceEnvParam(search string, replace string) {
	fileName := ".env"
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, search) {
			mask := fmt.Sprintf("%s[\\.A-Za-zА-Яа-я0-9]*", search)
			regex := regexp.MustCompile(mask)
			lines[i] = regex.ReplaceAllString(line, replace)
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0755)
	if err != nil {
		log.Fatalln(err)
	}
}

func PrepareEnv() {
	env := ".env"
	envExample := ".env.example"
	if !FileExists(env) && FileExists(envExample) {
		from, err := os.Open(envExample)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err := from.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		to, err := os.OpenFile(env, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err := to.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		_, err = io.Copy(to, from)
		if err != nil {
			log.Fatal(err)
		}
	}
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
