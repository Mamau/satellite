module satellite

go 1.14

require (
	github.com/creack/pty v1.1.11
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.9.0
	github.com/golang/mock v1.3.1
	github.com/gookit/color v1.3.8
	github.com/joho/godotenv v1.3.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.16.0 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/net v0.0.0-20210903162142-ad29c8ab022f // indirect
	golang.org/x/sys v0.0.0-20220730100132-1609e554cd39 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/mamau/satellite/internal => ./internal
