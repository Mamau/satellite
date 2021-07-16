package updater

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/gookit/color"

	"github.com/mamau/satellite/pkg"
)

const Version = "v0.15.1"

type SelfUpdater struct {
	Name           string
	CurrentVersion string
	LatestRelease  *Release
}

func NewSelfUpdater() *SelfUpdater {
	return &SelfUpdater{
		Name:           strings.Replace(os.Args[0], "./", "", -1),
		CurrentVersion: Version,
		LatestRelease:  fetchRelease(),
	}
}

func (s *SelfUpdater) Update() {

	if s.LatestRelease.Name == s.CurrentVersion {
		color.Cyan.Println("You have latest version")
		return
	}

	src := s.downloadLatest()

	data, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatalf("error while read source, err: %s\n", err)
	}

	if !s.replaceFile(data) {
		color.Red.Println("Error while updating")
	}

	color.Cyan.Println("Updating successfully")
}

func (s *SelfUpdater) replaceFile(fileData []byte) bool {
	dst := fmt.Sprintf("%s_old", s.Name)

	if err := ioutil.WriteFile(dst, fileData, 0744); err != nil {
		log.Fatalf("error while write file, err: %s\n", err)
		return false
	}

	if err := os.Rename(dst, s.Name); err != nil {
		log.Fatalf("error while rename file, err: %s\n", err)
		return false
	}

	return true
}

func (s *SelfUpdater) downloadLatest() string {
	currentOs := runtime.GOOS
	tmpFile := ""

	for _, v := range s.LatestRelease.Assets {
		if strings.Contains(v.Name, currentOs) {
			tmpFile = pkg.DownloadFile(v.Uri, s.Name, os.TempDir())
		}
	}
	return tmpFile
}
