package updater

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"

	"github.com/gookit/color"
)

func TestFetchRelease(t *testing.T) {
	data := []byte(`{"name":"someName","tag_name":"someTageName","assets":[{"name":"name1","browser_download_url":"uri1"},{"name":"name2","browser_download_url":"uri2"}]}`)
	var release Release
	if err := json.Unmarshal(data, &release); err != nil {
		color.Danger.Printf("cant unmarshal object, err: %s\n", err)
		os.Exit(1)
	}

	ctrl := gomock.NewController(GinkgoT())
	releaser := NewMockReleaser(ctrl)

	releaser.EXPECT().FetchRelease().Return(&release).MaxTimes(1)
	result := releaser.FetchRelease()
	assert.ObjectsAreEqual(release, result)
}
