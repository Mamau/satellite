package updater

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const source = "https://api.github.com/repos/Mamau/satellite/releases/latest"

type Asset struct {
	Name string `json:"name"`
	Uri  string `json:"browser_download_url"`
}

type Release struct {
	Name    string  `json:"name"`
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

func fetchRelease() *Release {
	res, err := http.Get(source)
	if err != nil {
		log.Fatalf("cant get info from github, err: %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("cant read body of response, err: %s\n", err)
	}

	var release Release
	if err := json.Unmarshal(body, &release); err != nil {
		log.Fatalf("cant unmarshal objecr, err: %s\n", err)
	}

	return &release
}
