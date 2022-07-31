package updater

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gookit/color"
)

const source = "https://api.github.com/repos/Mamau/satellite/releases/latest"

//go:generate mockgen -destination=releaser_mock.go -package=updater satellite/internal/updater Releaser

type Releaser interface {
	FetchRelease() *Release
}

type Asset struct {
	Name string `json:"name"`
	Uri  string `json:"browser_download_url"`
}

type Release struct {
	Name    string  `json:"name"`
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

func FetchRelease() *Release {
	res, err := http.Get(source)
	if err != nil {
		color.Danger.Printf("cant get info from github, err: %s\n", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		color.Danger.Printf("cant read body of response, err: %s\n", err)
		os.Exit(1)
	}

	var release Release
	if err := json.Unmarshal(body, &release); err != nil {
		color.Danger.Printf("cant unmarshal object, err: %s\n", err)
		os.Exit(1)
	}

	return &release
}
